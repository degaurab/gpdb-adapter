package config

import (
	"github.com/degaurab/gbdb-adapter/helper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	InstanceGroupName string    `yaml:"instance_group_name"`
	InstanceIP        string    `yaml:"instance_ip"`
	AdminUsername     string    `yaml:"admin_username"`
	AdminPassword     string    `yaml:"admin_password"`
	ConnectionPort    int       `yaml:"connection_port"`
	Templates         Templates `yaml:"templates"`
}

type Templates struct {
	BaseDir        string      `yaml:"basedir"`
	SchemaTemplate SQLTemplate `yaml:"schema_template"`
	UserTemplate   SQLTemplate `yaml:"user_template"`
}

type SQLTemplate struct {
	FileName string            `yaml:"file_name"`
	Vars     map[string]string `yaml:"vars"`
}

func LoadConfig(configFilePath string, logger *log.Logger) (Config, error) {
	config := Config{}

	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return config, helper.WrappedError("Error loading config file", err, logger)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, helper.WrappedError("Error parsing yaml file", err, logger)
	}

	logger.Println("DB config loaded:", config)
	//if config.InstanceIP == "" {
	//	return response.Config{}, helper.WrappedError("Error parsing yaml file", errors.New(""), logger)
	//}

	return config, nil
}
