package sdkfakes

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/s3"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/sanathkr/yaml"
)

type FakeOutput struct {
	Description string      `json:"Description"`
	Value       interface{} `json:"Value"`
	Ref         string      `json:"Ref"`
}

type FakeTemplate struct {
	Outputs map[string]FakeOutput `json:"Outputs"`
}

// NewHappyClient creates a fake AWSClient that attempts to stub the state
// transitions of a cloudformation stack for use in tests without real AWSClient
//
// The client will act like a "happy" aws, that transitions cleanly through:
//
// client.CreateStack()   -> ok
// client.DescribeStack() -> CREATE_IN_PROGRESS
// client.DescribeStack() -> CREATED_COMPLETED (few seconds later)
// client.UpdateStack()   -> ok (no update required)
// client.DeleteStack()   -> ok
// client.DescribeStack() -> DELETE_IN_PROGRESS
// client.DescribeStack() -> DELETE_COMPLETED (few seconds later)
//
// an attempt to call a client method outside of the happy path will cause error.
//
// stacks returned from DescribeStack methods will contain Outputs extracted from
// the template used in CreateStack.
func NewHappyClient(outputs map[string]string) *FakeClient {

	var transitionDelay = time.Second * 2 // how long before switching from CREATING->CREATED etc
	var client = &FakeClient{}
	var stack *cloudformation.Stack

	setStackState := func(s, reason string) {
		// initial/uncreated state
		if s == "" {
			stack = nil
			return
		}
		// set state
		stack.StackStatus = aws.String(s)
		stack.StackStatusReason = aws.String(reason)
	}

	setStackOutputsFromTemplate := func(templateYAML string) {
		var t FakeTemplate
		err := yaml.Unmarshal([]byte(templateYAML), &t)
		if err != nil {
			panic(err)
		}
		stack.Outputs = []*cloudformation.Output{}
		for k, v := range t.Outputs {
			stack.Outputs = append(stack.Outputs, &cloudformation.Output{
				Description: aws.String(v.Description),
				OutputKey:   aws.String(k),
				OutputValue: aws.String(outputs[k]),
			})
		}
	}

	client.DescribeStacksWithContextStub = func(context.Context, *cloudformation.DescribeStacksInput, ...request.Option) (*cloudformation.DescribeStacksOutput, error) {
		if stack == nil {
			return nil, ResourceNotFoundException
		}
		return &cloudformation.DescribeStacksOutput{
			Stacks: []*cloudformation.Stack{stack},
		}, nil
	}

	client.CreateStackWithContextStub = func(_ context.Context, input *cloudformation.CreateStackInput, o ...request.Option) (*cloudformation.CreateStackOutput, error) {
		if stack == nil {
			stack = &cloudformation.Stack{
				StackId:   aws.String(fmt.Sprintf("stack-%d", rand.Intn(10000))),
				StackName: input.StackName,
			}
			setStackState(cloudformation.StackStatusCreateInProgress, "fake-create-stack-called")
			// extract the cloudformation outputs from the given template
			if input.TemplateBody == nil {
				return nil, fmt.Errorf("TemplateBody is required")
			}
			// start timer to swtich to CREATE_COMPLETE state and add Outputs
			go func() {
				time.Sleep(transitionDelay)
				setStackOutputsFromTemplate(*input.TemplateBody)
				setStackState(cloudformation.StackStatusCreateComplete, "fake-creation-timer-completed")
			}()
			return &cloudformation.CreateStackOutput{
				StackId: stack.StackId,
			}, nil
		}
		return nil, fmt.Errorf("CANNOT_CREATE_ALREADY_CREATED")
	}

	client.DeleteStackWithContextStub = func(context.Context, *cloudformation.DeleteStackInput, ...request.Option) (*cloudformation.DeleteStackOutput, error) {
		if stack == nil {
			return nil, fmt.Errorf("CANNOT_DELETE_BEFORE_CREATE")
		}
		switch *stack.StackStatus {
		case cloudformation.StackStatusCreateComplete, cloudformation.StackStatusUpdateComplete, cloudformation.StackStatusUpdateRollbackComplete, cloudformation.StackStatusRollbackComplete:
			go func() {
				// after a while transition to DELETE_COMPLETE state
				time.Sleep(transitionDelay)
				setStackState(cloudformation.StackStatusDeleteComplete, "fake-deletion-timer-completed")
			}()
			return &cloudformation.DeleteStackOutput{}, nil
		default:
			return nil, fmt.Errorf("CANNOT_DELETE_FROM_CURRENT_STATE: %s", *stack.StackStatus)
		}
	}

	client.UpdateStackWithContextStub = func(context.Context, *cloudformation.UpdateStackInput, ...request.Option) (*cloudformation.UpdateStackOutput, error) {
		return nil, NoUpdateRequiredException
	}

	client.AssumeRoleReturns(client)

	getRoleCredsLater := time.After(time.Second * 60)
	client.GetRoleCredentialsStub = func(string) *credentials.Credentials {
		select {
		case <-getRoleCredsLater:
			return credentials.NewStaticCredentialsFromCreds(credentials.Value{
				AccessKeyID: "some second value",
				SecretAccessKey: "some new secret!",
				SessionToken: "different session token",
				ProviderName: "static but really just a fake to make client happy",
			})
		default:
			return credentials.NewStaticCredentialsFromCreds(credentials.Value{
				AccessKeyID: "some value",
				SecretAccessKey: "zomg seekrits",
				SessionToken: "session here",
				ProviderName: "static but really just a fake to make client happy",
			})
		}
	}

	getAuthTokenLater := time.After(time.Second * 60)
	client.GetAuthorizationTokenWithContextStub = func(context.Context, *ecr.GetAuthorizationTokenInput, ...request.Option) (*ecr.GetAuthorizationTokenOutput, error) {
		select {
		case <-getAuthTokenLater:
			return &ecr.GetAuthorizationTokenOutput{
				AuthorizationData: []*ecr.AuthorizationData{
					{
						AuthorizationToken: aws.String("d2hhdGV2ZXJ0aHJlZTp3aGF0ZXZlcmZvdXI="), // whateverthree:whateverfour
						ExpiresAt:          nil,
						ProxyEndpoint:      aws.String("https://011571571136.dkr.ecr.eu-west-2.amazonaws.com"),
					},
				},
			}, nil
		default:
			return &ecr.GetAuthorizationTokenOutput{
				AuthorizationData: []*ecr.AuthorizationData{
					{
						AuthorizationToken: aws.String("d2hhdGV2ZXJvbmU6d2hhdGV2ZXJ0d28="), // whateverone:whatevertwo
						ExpiresAt:          nil,
						ProxyEndpoint:      aws.String("https://011571571136.dkr.ecr.eu-west-2.amazonaws.com"),
					},
				},
			}, nil
		}
	}

	client.DescribeImagesPagesWithContextStub = func(_ context.Context, input *ecr.DescribeImagesInput, fn func(page *ecr.DescribeImagesOutput, lastPage bool) bool, o ...request.Option) error {
		fn(
			&ecr.DescribeImagesOutput{
				ImageDetails: []*ecr.ImageDetail{
					&ecr.ImageDetail{
						ImageDigest:      aws.String("sha256:some long sha256 sum"),
						ImagePushedAt:    aws.Time(time.Now()),
						ImageScanStatus:  &ecr.ImageScanStatus{
							Description:  aws.String("not done"),
							Status:       aws.String("fake client happy"),
						},
						ImageSizeInBytes: aws.Int64(42),
						ImageTags:        []*string{
							aws.String("latest"),
						},
						RegistryId:       input.RegistryId,
						RepositoryName:   input.RepositoryName,
					},
				},
			},
			true,
		)
		return nil
	}
	client.BatchDeleteImageWithContextStub = func(_ context.Context, input *ecr.BatchDeleteImageInput, o ...request.Option) (*ecr.BatchDeleteImageOutput, error) {
		return &ecr.BatchDeleteImageOutput{
			Failures: []*ecr.ImageFailure{},
			ImageIds: input.ImageIds,
		}, nil
	}

	client.ListObjectsV2PagesWithContextStub = func(_ context.Context, _ *s3.ListObjectsV2Input, fn func(page *s3.ListObjectsV2Output, lastPage bool) bool, o ...request.Option) error {
		fn(
			&s3.ListObjectsV2Output{
				Contents: []*s3.Object{
					&s3.Object{
						ETag:         aws.String("some etag"),
						Key:          aws.String("important.file"),
						LastModified: aws.Time(time.Now()),
						Owner:        &s3.Owner{
							DisplayName: aws.String("Alex"),
							ID:          aws.String("0"),
						},
						Size:         aws.Int64(42),
						StorageClass: aws.String("STANDARD"),
					},
				},
			},
			true,
		)
		return nil
	}
	client.DeleteObjectsWithContextStub = func(context.Context, *s3.DeleteObjectsInput, ...request.Option) (*s3.DeleteObjectsOutput, error) {
		return &s3.DeleteObjectsOutput{
			Errors: []*s3.Error{},
		}, nil
	}

	return client
}
