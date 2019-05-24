package config

import (
	"github.com/degaurab/gbdb-adapter/helper"
	"github.com/degaurab/gbdb-adapter/response"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func LoadCatalog(catalogFilePath string, logger *log.Logger) (response.CatalogResponse, error) {
	catalog := response.CatalogResponse{}
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
