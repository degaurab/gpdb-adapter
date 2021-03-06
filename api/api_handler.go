package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/degaurab/gpdb-adapter/config"
	gpdb_client "github.com/degaurab/gpdb-adapter/gpdb-client"
	"github.com/degaurab/gpdb-adapter/helper"
	"github.com/gorilla/mux"
)

func NewApiHandler(log *log.Logger, confPath string, catPath string, r *mux.Router) {
	api := ApiHandler{
		logger:      log,
		configPath:  confPath,
		catalogPath: catPath,
	}
	r.HandleFunc("/v1/catalog", api.serviceCatalog).Methods("GET")
	r.HandleFunc("/v1/create_binding", api.createBinding).Methods("PUT")
	r.HandleFunc("/v1/delete_binding/{binding_id}", api.deleteBinding).Methods("PUT")
}

type ApiHandler struct {
	logger      *log.Logger
	configPath  string
	catalogPath string
}

type response struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
}

func (api ApiHandler) serviceCatalog(httpWriter http.ResponseWriter, httpReqest *http.Request) {
	api.logger.Println("loding config")
	resp := response{}

	catalog, err := config.LoadCatalog(api.catalogPath, api.logger)
	if err != nil {
		resp.Error = err.Error()
		api.respond(httpWriter, 500, resp)
		return
	}

	resp.Result = catalog
	api.logger.Println("response:", resp)

	api.respond(httpWriter, 200, resp)
}

func (api ApiHandler) createBinding(httpWriter http.ResponseWriter, httpReqest *http.Request) {
	//return nil, nil
	api.logger.Println("Started Creating Binding")
	resp := response{}

	/*
		TODO: get the requested username and db from data payload
	*/
	//data := httpReqest.Body

	// TODO: loading confing on each request:
	// - is costly
	// - but we dont have to reload app for config changes
	// - security issues: config can be change without restarting app
	//
	c, err := config.LoadConfig(api.configPath, api.logger)
	if err != nil {
		resp.Error = err.Error()
		api.respond(httpWriter, 500, resp)
		return
	}

	driver := gpdb_client.DBDriver{
		User:       c.AdminUsername,
		Password:   c.AdminPassword,
		Port:       c.ConnectionPort,
		Hostname:   c.InstanceIP,
		DBTemplate: c.Templates,
	}

	api.logger.Println(driver)

	//Creating DB, User and grant access
	randUsername := helper.RandStringBytes(5)
	dbName := "client" + randUsername
	user, err := driver.InitializeDBForUser(dbName, randUsername, api.logger)
	if err != nil {
		resp.Error = err.Error()
		api.respond(httpWriter, 500, resp)
		return
	}

	resp.Result = user
	api.respond(httpWriter, 200, resp)
}

func (api ApiHandler) deleteBinding(httpWriter http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)
	bindingID := data["binding_id"]

	resp := response{}

	c, err := config.LoadConfig(api.configPath, api.logger)
	if err != nil {
		resp.Error = err.Error()
		api.respond(httpWriter, 500, resp)
		return
	}

	driver := gpdb_client.DBDriver{
		User:     c.AdminUsername,
		Password: c.AdminPassword,
		Port:     c.ConnectionPort,
		Hostname: c.InstanceIP,
	}

	//TODO: extract from payload
	driver.DeleteDatabase(bindingID, api.logger)
	api.respond(httpWriter, 200, resp)
}

func (api ApiHandler) respond(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		api.logger.Println(err, "encoding response")
	}
}
