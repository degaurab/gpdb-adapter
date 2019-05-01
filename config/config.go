package config

import (
	"github.com/degaurab/gbdb-adapter/helper"
	"github.com/degaurab/gbdb-adapter/response"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)


func LoadConfig(configFilePath string, logger *log.Logger) (response.Config, error) {
	config := response.Config{}

	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return config, helper.WrappedError("Error loading config file", err, logger)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, helper.WrappedError("Error parsing yaml file", err, logger)
	}

	return config, nil
}
