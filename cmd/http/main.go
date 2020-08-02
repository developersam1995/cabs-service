package main

import (
	cs "github.com/developersam1995/cabs-service/lib/cab-service"
	"github.com/developersam1995/cabs-service/storage"
	"github.com/gin-gonic/gin"
)

func setupRouter(b cs.BookingService) *gin.Engine {
	r := gin.Default()
	r.POST("/booking", MakeBookingsPoster(b))
	return r
}

func main() {
	db := storage.New(storage.DbConfig{})
	bookingService := cs.NewBookingService(db)
	r := setupRouter(bookingService)
	r.Run(":8080")
}
