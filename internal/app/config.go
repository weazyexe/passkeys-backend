package app

import (
	"gopkg.in/yaml.v2"
	"os"
)

type PasskeysDemoConfig struct {
	Webauthn Webauthn `yaml:"webauthn"`
}

type Webauthn struct {
	RPDisplayName string   `yaml:"rpDisplayName"`
	RPID          string   `yaml:"rpId"`
	RPOrigins     []string `yaml:"rpOrigins"`
}

func ReadConfig(path string) (*PasskeysDemoConfig, error) {
	configYaml, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := PasskeysDemoConfig{}
	if err := yaml.Unmarshal(configYaml, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
