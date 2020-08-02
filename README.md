## Cabs services.

* cmd/http contains the main http service package
* lib/cabs-service contains the core business logic of the server
* storage is self explanatory

> `config_example.json` in `cmd/http` package shows an example of server config. Please copy this file as `config.json` and modify. The build takes the location of this file as the first argument. If not provided, assumes default values.

> Testing of controllers expects a local mysql with the following values, which can be changed with a config.json file 
* db : cabs_db,
* user : server_x
* password: testpassword

TODO:
To confirm booking 
To register users
To authenticate and authorize
To register cabs