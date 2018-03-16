package parser

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type Payload struct {
	Machine struct {
		Environment map[string]string `yaml:"environment"`
	} `yaml:"machine"`
	Test struct {
		Override []string `yaml:"override"`
	} `yaml:"test"`
}

func NewPayloadFromFilePath(path string) (*Payload, error) {
	var payload Payload
	config, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(config, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}
