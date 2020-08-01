package main

import (
	"net/http"

	cs "github.com/developersam1995/cabs-service/lib/cab-service"
	"github.com/gin-gonic/gin"
)

type booker interface {
	BookCab(cs.BookRequest) (bookingID int, err error)
}

type location struct {
	Lat float64 `json:"lat" binding:"required"`
	Lon float64 `json:"lon" binding:"required"`
}

type bookingPost struct {
	From       location `json:"from" binding:"required"`
	To         location `json:"to" binding:"required"`
	UserID     int      `json:"user_id" binding:"required"`
	PickupTime int      `json:"pickup_time" binding:"required"`
}

func MakeBookingsPoster(b booker) gin.HandlerFunc {
	return func(c *gin.Context) {

		reqBody := bookingPost{}
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		br := cs.BookRequest{
			From: cs.Location{
				Lat: reqBody.From.Lat,
				Lon: reqBody.From.Lon,
			},
			To: cs.Location{
				Lat: reqBody.To.Lat,
				Lon: reqBody.To.Lon,
			},
			UserID:     reqBody.UserID,
			PickupTime: reqBody.PickupTime,
		}
		id, err := b.BookCab(br)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": serverErrMsg,
			})
			return
		}
		c.JSON(
			http.StatusCreated,
			gin.H{
				"booking_id": id,
			},
		)
	}
}
