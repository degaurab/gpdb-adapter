package main

import (
	"os"

	//"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/degaurab/gbdb-adapter/api"
)

const CatalogPath = "/tmp/catalog.yml"
const ConfigPath = "/tmp/service-config.yml"


func main() {

	log := log.New(os.Stderr,"[gbpd-service-adapter]", log.LstdFlags)
	r := mux.NewRouter()

	api.NewApiHandler(log, ConfigPath, CatalogPath, r)


	log.Println("Server Started. Listning on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}