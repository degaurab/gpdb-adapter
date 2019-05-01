package config

import (
	"github.com/degaurab/gbdb-adapter/helper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Catalog struct {
	Services []Service
}

type Service struct {
	Name string `yaml:"name"`
	Id string `yaml:"id"`
	Desc string `yaml:"description"`
	Plans []Plan `yaml:"plans"`
}


type Plan struct {
	Name string `yaml:"name"`
	Id string `yaml:"id"`
	Metadata map[string]interface{} `yaml:"metadata"`
}



func LoadCatalog(catalogFilePath string, logger *log.Logger) (Catalog, error) {
	catalog := Catalog{}

	yamlFile, err := ioutil.ReadFile(catalogFilePath)
	if err != nil {
		return catalog, helper.WrappedError("Error loading catalog file", err, logger)
	}
	
	err = yaml.Unmarshal(yamlFile, &catalog)
	if err != nil {
		return catalog, helper.WrappedError("Error parsing yaml file", err, logger)
	}

	return catalog, nil
}
