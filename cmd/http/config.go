package main

type dbConfig struct {
	Host        string `json:"host"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Db          string `json:"db"`
	Pooling     bool   `json:"pooling"`
	Connections int    `json:"connections"`
}
type configuration struct {
	Port          int `json:"port"`
	Db            dbConfig
	MaxCabGetDist int `json:"max_cab_get_dist"` // See the Get cabs api
}
