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

package pipeline

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestEvaluatePipeline(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

  testYaml := `git_source: &git_source
  uri: https://github.com/concourse/concourse
  branch: master
  extra: ((var))

resources:
  - name: repo
    type: git
    source:
      <<: *git_source
      branch: test

jobs:
  - name: test
    plan:
      - get: repo`

  expectedYaml := []byte(`git_source:
  branch: master
  extra: ((var))
  uri: https://github.com/concourse/concourse
jobs:
- name: test
  plan:
  - get: repo
resources:
- name: repo
  source:
    branch: test
    extra: ((var))
    uri: https://github.com/concourse/concourse
  type: git`)

  actualYaml, err := EvaluatePipeline(testYaml)

  g.Expect(err).To(Equal(nil), fmt.Sprintf("got error: %s", err))
  g.Expect(actualYaml).To(Equal(expectedYaml))
}
