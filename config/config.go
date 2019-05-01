package config

import (
	"github.com/degaurab/gbdb-adapter/helper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	GPDBInstanceGroupName string `yaml:"gpdb_instance_group_name"`
	GPDBInstanceIP string `yaml:"gpdb_instance_ip"`
	GPDBAdminUsername string `yaml:"gpdb_admin_username"`
	GPDBAdminPassword string `yaml:"gpdb_admin_password"`
	ConnectionPort int `yaml:"connection_port"`
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

	return config, nil
}
