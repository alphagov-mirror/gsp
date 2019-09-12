/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	accessv1beta1 "github.com/alphagov/gsp/components/service-operator/apis/access/v1beta1"
	databasev1beta1 "github.com/alphagov/gsp/components/service-operator/apis/database/v1beta1"
	queuev1beta1 "github.com/alphagov/gsp/components/service-operator/apis/queue/v1beta1"
	"github.com/alphagov/gsp/components/service-operator/controllers"
	"github.com/alphagov/gsp/components/service-operator/internal/aws/sdk"
	"github.com/alphagov/gsp/components/service-operator/internal/aws/sdk/sdkfakes"
	"github.com/alphagov/gsp/components/service-operator/internal/env"
	core "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t,
		"Controller Suite",
		[]Reporter{envtest.NewlineReporter{}})
}

// SetupControllerEnv will create a real kubernetes control plane, setup a
// manager and aws client then call controllerFn which is expected to return
// the controller under test.  SetupControllerEnv will return a wrapped version
// of the controller which can be used to inspect Reconcile errors and a
// teardown function that should be called after the test is complete.
// It is probably not practical to run this in parallel
func SetupControllerEnv() (client.Client, func()) {
	os.Setenv("CLOUD_PROVIDER", "aws")
	os.Setenv("CLUSTER_NAME", "xxx")
	ctx := context.Background()

	log := zap.LoggerTo(GinkgoWriter, false)
	logf.SetLogger(log)

	testEnv := &envtest.Environment{
		CRDDirectoryPaths: []string{filepath.Join("..", "config", "crd", "bases")},
	}

	var err error
	cfg, err := testEnv.Start()
	Expect(err).ToNot(HaveOccurred())
	Expect(cfg).ToNot(BeNil())

	err = databasev1beta1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = queuev1beta1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = accessv1beta1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme: scheme.Scheme,
	})
	Expect(err).ToNot(HaveOccurred())

	k8sClient := mgr.GetClient()
	Expect(k8sClient).ToNot(BeNil())

	// wait for control plane to be happy
	Eventually(func() error {
		return k8sClient.List(ctx, &core.SecretList{})
	}, time.Second*20).Should(Succeed())

	// create channel for teardown
	mgrStopChan := make(chan struct{})

	// controllers under test
	cs := []controllers.Controller{
		controllers.SQSCloudFormationController(newAWSClient()),
		controllers.PrincipalCloudFormationController(newAWSClient()),
		controllers.PostgresCloudFormationController(newAWSClient()),
	}

	// wrap controllers in error checers and register with manager
	for i := range cs {
		controller := &controllers.ControllerWrapper{
			Reconciler: cs[i],
		}
		err = controller.SetupWithManager(mgr)
		Expect(err).ToNot(HaveOccurred())
		go reconcileErrorMonitor(controller, mgrStopChan)
	}

	By("starting controller manager")
	go func() {
		err = mgr.Start(mgrStopChan)
		Expect(err).ToNot(HaveOccurred())
	}()

	return mgr.GetClient(), func() {
		By("stopping controller manager")
		close(mgrStopChan)
		By("stopping control plane")
		Expect(testEnv.Stop()).To(Succeed())
	}
}

func reconcileErrorMonitor(controller *controllers.ControllerWrapper, stop chan struct{}) {
	defer GinkgoRecover()
	// fail test if we see any reconcile errors
	for {
		select {
		case <-stop:
			return
		case <-time.After(time.Millisecond * 250):
			Expect(controller.Err()).ToNot(HaveOccurred())
		}
	}
}

func newAWSClient() sdk.Client {
	if env.AWSIntegrationTestEnabled() {
		return sdk.NewClient()
	} else {
		// set dummy values when running against mock
		os.Setenv("AWS_RDS_SECURITY_GROUP_ID", "dummy-value")
		os.Setenv("AWS_RDS_SUBNET_GROUP_NAME", "dummy-value")
		os.Setenv("AWS_PRINCIPAL_SERVER_ROLE_ARN", "dummy-value")
		os.Setenv("AWS_PRINCIPAL_PERMISSIONS_BOUNDARY_ARN", "dummy-value")
		return sdkfakes.NewHappyClient()
	}
}
