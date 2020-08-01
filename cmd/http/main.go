package main

import (
	"github.com/developersam1995/cabs-service/storage"
	"github.com/gin-gonic/gin"
)

func setupRouter(db *storage.Db) *gin.Engine {
	r := gin.Default()
	r.POST("/booking", MakeBookingsPoster(db))
	return r
}

func main() {
	db := storage.New(storage.DbConfig{})
	r := setupRouter(db)
	r.Run(":8080")
}
