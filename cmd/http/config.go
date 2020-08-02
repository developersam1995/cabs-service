package main

import "github.com/developersam1995/cabs-service/storage"

type configuration struct {
	Port          int `json:"port"`
	Db            storage.DbConfig
	MaxCabGetDist int `json:"max_cab_get_dist"` // See the Get cabs api
}
