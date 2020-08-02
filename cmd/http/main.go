package main

import (
	cs "github.com/developersam1995/cabs-service/lib/cab-service"
	"github.com/developersam1995/cabs-service/storage"
	"github.com/gin-gonic/gin"
)

type jsonMock map[string]interface{}

func setupBookingRoutes(r *gin.Engine, b cs.BookingService) {
	r.POST("/booking", MakeBookingsPoster(b))
	r.GET("/booking/:userID", MakeBookingsGetter(b))
}

func setupCabRoutes(r *gin.Engine, c cs.CabsService) {
	r.GET("/cabs", MakeCabsGetter(c))
}

func main() {
	r := gin.Default()
	db := storage.New(storage.DbConfig{})

	bookingService := cs.NewBookingService(db)
	cabsService := cs.CabsService(db)
	setupBookingRoutes(r, bookingService)
	setupCabRoutes(r, cabsService)
	r.Run(":8080")
}
