## Cabs services.

* cmd/http contains the main http service package
* lib/cabs-service contains the core business logic of the server
* storage is self explanatory

> `config_example.json` shows an example of server config. Please copy this file as `config.json` and modify. The build takes the location of this file as the first argument. If not provided, assumes default values.

* To build run `make build` it will output to bin folder
* To test run `make test`
* For swagger docs open `{host}/swagger/index.html` in your browser.

TODO:
* To confirm booking 
* To register users
* To authenticate and authorize
* To register cabs