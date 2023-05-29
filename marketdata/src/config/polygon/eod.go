package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type EODConfig struct {
	Symbols  []string `yaml:"symbols"`
	Date     string   `yaml:"date"`
	Adjusted bool     `yaml:"adjusted"`
}

func NewEODConfig(fpath string) (EODConfig, error) {
	config, err := parseEODConfig(fpath)
	if err != nil {
		return EODConfig{}, fmt.Errorf("%w: failed to parse config from file: filepath=%s", err, fpath)
	}
	return config, nil
}

func parseEODConfig(fpath string) (EODConfig, error) {
	content, err := ioutil.ReadFile(fpath)
	if err != nil {
		return EODConfig{}, err
	}

	var config EODConfig
	if err := yaml.Unmarshal(content, &config); err != nil {
		return EODConfig{}, err
	}
	return config, nil
}
