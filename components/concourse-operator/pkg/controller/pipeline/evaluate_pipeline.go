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
//	"fmt"
//  "io/ioutil"
//	"os"

	"github.com/concourse/concourse/vars"
)

func EvaluatePipeline(pipelineYAML string) ([]byte, error) {
  // run pipelineYAML through fly
  // -- write to temp file
  /*
  tmpfile, err := ioutil.TempFile("", fmt.Sprintf("%s_%s_*", teamName, pipelineName))
  if err != nil {
    return nil, err
  }

  defer os.Remove(tmpfile.Name())

  if _, err := tmpfile.Write([]byte(pipelineYAML)); err != nil {
    return nil, err
  }

  if err := tmpfile.Close(); err != nil {
    return nil, err
  }

  // -- build new YamlTemplate
  yamlTemplate := .NewYamlTemplateWithParams(atc.PathFlag(tmpfile.Name()), nil, nil, nil)

  */
  // -- use evaluated yaml string
  evaluatedYaml, err := vars.NewTemplateResolver([]byte(pipelineYAML), nil).Resolve(false, false)
  if err != nil {
    return nil, err
  }

  return evaluatedYaml, nil
}
