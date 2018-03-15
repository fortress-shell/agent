package parser

import (
	"os"
	"io/ioutil"
	"github.com/fortress-shell/agent/steps"
	"github.com/go-yaml/yaml"
)

type Payload struct {
	Machine struct {
		Environment map[string]string `yaml:"environment"`
	} `yaml:"machine"`
	// Dependencies struct {
	// 	Override []string `yaml:"override"`
	// } `yaml:"dependencies"`
	Test struct {
		Override []string `yaml:"override"`
	} `yaml:"test"`
	// Deployment struct {
	// 	Override []string `yaml:"override"`
	// }
}

func NewPayloadFromFilePath(path string) (*Payload, error) {
	var payload Payload
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()
	byteValue, _ := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(byteValue, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func GenerateSteps(p *Payload) []steps.Step {
	var tasks []steps.Step
	for _, v := range p.Test.Override {
		build := steps.OverrideBuildStep{
			Command:     v,
			Environment: p.Machine.Environment,
		}
		tasks = append(tasks, &build)
	}
	return tasks
}
