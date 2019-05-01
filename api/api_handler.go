package api

import (
	"encoding/json"
	"github.com/degaurab/gbdb-adapter/config"
	"github.com/degaurab/gbdb-adapter/gpdb-client"
	"github.com/degaurab/gbdb-adapter/helper"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func NewApiHandler(log *log.Logger, confPath string, catPath string, r *mux.Router) {
	api := ApiHandler{
		logger: log,
		configPath: confPath,
		catalogPath: catPath,
	}
	r.HandleFunc("/v2/catalog", api.ServiceCatalog).Methods("GET")
	r.HandleFunc("/v2/create_binding", api.CreateBinding).Methods("POST")
	r.HandleFunc("/v2/delete_binding", api.DeleteBinding).Methods("POST")
}

type ApiHandler struct {
	logger *log.Logger
	configPath string
	catalogPath string
}


type response struct {
	result interface{} `json:"result"`
	error string `json:"error"`
}

func (api ApiHandler) ServiceCatalog(httpWriter http.ResponseWriter, httpReqest *http.Request)  {
	api.logger.Println("loding config")
	resp := response{}

	catalog, err := config.LoadCatalog(api.catalogPath, api.logger)
	if err != nil {
		resp.error = err.Error()
	}

	resp.result = catalog
	api.logger.Println("response:", resp)

	api.respond(httpWriter, 200, resp)
}

func (api ApiHandler) CreateBinding(httpWriter http.ResponseWriter, httpReqest *http.Request) {
	//return nil, nil
	api.logger.Println("Started Creating Binding")
	resp := response{}

	/*
	TODO: get the requested username and db from data payload
	 */
	//data := httpReqest.Body

	c, err := config.LoadConfig(api.configPath, api.logger)
	if err != nil {
		resp.error = err.Error()
		json.NewEncoder(httpWriter).Encode(resp)
	}

	driver := gpdb_client.DBDriver{
		User: c.GPDBAdminUsername,
		Password: c.GPDBAdminPassword,
		Port: c.ConnectionPort,
		Hostname: c.GPDBInstanceIP,
	}

	//Creating DB, User and grant access
	randUsername := helper.RandStringBytes(5)
	dbName := "client-" + randUsername
	user, err := driver.InitializeDBForUser(dbName, randUsername, api.logger)
	if err != nil {
		resp.error = err.Error()
	}

	resp.result = user

	api.logger.Println(resp)
	api.respond(httpWriter, 200, resp)
}


func (api ApiHandler) DeleteBinding(httpWriter http.ResponseWriter, r *http.Request)  {
	resp := response{}

	c, err := config.LoadConfig(api.configPath, api.logger)
	if err != nil {
		resp.error = err.Error()
		json.NewEncoder(httpWriter).Encode(resp)
	}

	driver := gpdb_client.DBDriver{
		User: c.GPDBAdminUsername,
		Password: c.GPDBAdminPassword,
		Port: c.ConnectionPort,
		Hostname: c.GPDBInstanceIP,
	}

	//TODO: extract from payload
	driver.DeleteDatabase("some-exisiting-DB")
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
	output,_ := json.Marshal(response)
	io.WriteString(w, string(output))
}

