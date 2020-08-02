package main

import (
	cs "github.com/developersam1995/cabs-service/lib/cab-service"
	"github.com/developersam1995/cabs-service/storage"
	"github.com/gin-gonic/gin"
)

type jsonMock map[string]interface{}

func setupBookingRoutes(r *gin.Engine, b cs.BookingService) {
	r.POST("/booking", makeBookingsPoster(b))
	r.GET("/booking/:userID", makeBookingsGetter(b))
}

func setupCabRoutes(r *gin.Engine, c cs.CabsService) {
	r.GET("/cabs", makeCabsGetter(c))
}

func main() {
	r := gin.Default()
	db := storage.New(storage.DbConfig{})

	bookingService := cs.NewBookingService(db)
	cabsService := cs.NewCabsService(db)

	setupBookingRoutes(r, bookingService)
	setupCabRoutes(r, cabsService)

	r.Run(":8080")
}
