## POC: GBDB Adapter

- Designed based on `openservicebroker` api
- No tests 



### TODO:

- update API endpoints to match brokerAPI endpoints

    ```$xslt
	router.HandleFunc("/v2/service_instances/{instance_id}", handler.getInstance).Methods("GET")
	router.HandleFunc("/v2/service_instances/{instance_id}", handler.provision).Methods("PUT")
	router.HandleFunc("/v2/service_instances/{instance_id}", handler.deprovision).Methods("DELETE")
	router.HandleFunc("/v2/service_instances/{instance_id}/last_operation", handler.lastOperation).Methods("GET")
	router.HandleFunc("/v2/service_instances/{instance_id}", handler.update).Methods("PATCH")

	router.HandleFunc("/v2/service_instances/{instance_id}/service_bindings/{binding_id}", handler.getBinding).Methods("GET")
	router.HandleFunc("/v2/service_instances/{instance_id}/service_bindings/{binding_id}", handler.bind).Methods("PUT")
	router.HandleFunc("/v2/service_instances/{instance_id}/service_bindings/{binding_id}", handler.unbind).Methods("DELETE")

	router.HandleFunc("/v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation", handler.lastBindingOperation).Methods("GET")
    ```


- Fix scope of `GRANT PRIVILAGES` command
 
```psql
postgres=# \l
                                   List of databases
    Name     |  Owner   | Encoding |   Collate   |    Ctype    |   Access privileges
-------------+----------+----------+-------------+-------------+-----------------------
 clientxvlbz | postgres | UTF8     | en_CA.UTF-8 | en_CA.UTF-8 | =Tc/postgres         +
             |          |          |             |             | postgres=CTc/postgres+
             |          |          |             |             | xvlbz=CTc/postgres
 gbb         | postgres | UTF8     | en_CA.UTF-8 | en_CA.UTF-8 |
 postgres    | postgres | UTF8     | en_CA.UTF-8 | en_CA.UTF-8 |
 template0   | postgres | UTF8     | en_CA.UTF-8 | en_CA.UTF-8 | =c/postgres          +
             |          |          |             |             | postgres=CTc/postgres
 template1   | postgres | UTF8     | en_CA.UTF-8 | en_CA.UTF-8 | =c/postgres          +
             |          |          |             |             | postgres=CTc/postgres
(5 rows)

```

- Make sure delete works.
- Move testing to container



### Build and Run

```bash
#-----------------------
$ cd github.com/degaurab/gbdp-adapter
$ go build main.go

## run
$ ./main

## Add catalog and config file to `/tmp`
## Samples are in `config/samples`
## Default path:
### //CatalogPath for testing
### const CatalogPath = "/tmp/catalog.yml"
### //ConfigPath for testing
### const ConfigPath = "/tmp/service-config.yml"
# -----------------------


### different terminal

## to view catalog
$ curl -X GET http://localhost:8080/v2/catalog | jq .

## to create binding (need to be fixed)
$ curl -X PUT http://localhost:8080/v2/create_binding | jq .

```

