package parser

import (
	"github.com/fortress-shell/agent/steps"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
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

type GeneratedSteps struct {
	Checkout      []steps.Step
	OverrideBuild []steps.Step
}

func GenerateSteps(p *Payload) *GeneratedSteps {
	tasks := &GeneratedSteps{}
	checkout := &steps.OverrideCheckoutStep{
		Environment: p.Machine.Environment,
	}
	tasks.Checkout = append(tasks.Checkout, checkout)
	for _, v := range p.Test.Override {
		build := &steps.OverrideBuildStep{
			Command:     v,
			Environment: p.Machine.Environment,
		}
		tasks.OverrideBuild = append(tasks.OverrideBuild, build)
	}
	return tasks
}
