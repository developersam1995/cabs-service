## Cabs services.

* cmd/http contains the main http service package
* lib/cabs-service contains the core logic of the server
* storage is self explanatory

> `config_example.json` in `cmd/http` package shows an example of server config. Please copy this file as `config.json` and modify. The build takes the location of this file as the first argument. If not provided, assumes default values.