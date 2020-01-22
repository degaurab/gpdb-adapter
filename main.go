package main

import (
	"flag"
	"os"

	//"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/degaurab/gpdb-adapter/api"
)

////CatalogPath for testing
//const CatalogPath = "/tmp/catalog.yml"
////ConfigPath for testing
//const ConfigPath = "/tmp/service-config.yml"
//

func main() {

	log := log.New(os.Stderr, "[gbpd-service-adapter]", log.LstdFlags)
	configPath := flag.String(
		"config",
		"/tmp/service-config.yml",
		"Path for service-config.yml",
	)
	catalogPath := flag.String(
		"catalog",
		"/tmp/catalog.yml",
		"Path for catalog.yml",
	)

	flag.Parse()
	log.Println("ConfigPath:", *configPath, "::", "CatalogPath:", *catalogPath)

	r := mux.NewRouter()

	api.NewApiHandler(log, *configPath, *catalogPath, r)

	log.Println("Server Started. Listning on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
