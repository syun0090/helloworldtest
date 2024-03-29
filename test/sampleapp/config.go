/*
Copyright 2019 The Knative Authors

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

package sampleapp

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

const (
	// using these defaults if not provided, see useDefaultIfNotProvided function below
	defaultSrcDir = "../../docs/serving/samples/hello-world/helloworld-%s"

	// ActionMsg serves as documentation purpose, which will be referenced for
	// clearly displaying error messages.
	ActionMsg = "All files required for running sample apps are checked " +
		"against README.md, the content of source files should be identical with what's " +
		"in README.md file, the list of the files to be verified is the same set of files " +
		"used for running sample apps, they are configured in `/test/sampleapp/config.yaml`. " +
		"If an exception is needed the file can be configured to be copied as a separate step " +
		"in `PreCommand` such as: " +
		"https://github.com/knative/docs/blob/65f7b402fee7f94dfbd9e4512ef3beed7b85de66/test/sampleapp/config.yaml#L4"
)

// AllConfigs contains all LanguageConfig
type AllConfigs struct {
	Languages []LanguageConfig `yaml:"languages"`
}

// LanguageConfig contains all information for building/deploying an app
type LanguageConfig struct {
	Language string   `yaml:"language"`
	SrcDir   string   `yaml:"srcDir"` // Directory contains sample code
	Copies   []string `yaml:"copies"` // Files to be copied by the user from SrcDir
}

// UseDefaultIfNotProvided sets default value of SrcDir if not provided
func (lc *LanguageConfig) UseDefaultIfNotProvided() {
	if "" == lc.SrcDir {
		lc.SrcDir = fmt.Sprintf(defaultSrcDir, lc.Language)
	}
}

// GetConfigs parses a config yaml file and return AllConfigs struct
func GetConfigs(configPath string) (AllConfigs, error) {
	var lcs AllConfigs
	content, err := ioutil.ReadFile(configPath)
	if nil == err {
		err = yaml.Unmarshal(content, &lcs)
	}
	return lcs, err
}
