package main

import (
	"net/http"

	cs "github.com/developersam1995/cabs-service/lib/cab-service"
	"github.com/gin-gonic/gin"
)

type BookingReq struct {
	FromLat    float64 `json:"from_lat" binding:"required"`
	FromLon    float64 `json:"from_lon" binding:"required"`
	ToLat      float64 `json:"to_lat" binding:"required"`
	ToLon      float64 `json:"to_lon" binding:"required"`
	UserID     int     `json:"user_id" binding:"required"`
	PickupTime int     `json:"pickup_time" binding:"required"`
}

type BookingID struct {
	id int `json:"booking_id"`
}

// @Summary Request for a cab booking
// @Produce  json
// @Accept   json
// @Success 200 {object} BookingID
// @Failure 400 {object} ErrResp
// @Param body body BookingReq true "Booking request"
// @Router /booking [post]
func makeBookingsPoster(b cs.BookingService) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Should use userID from an authenticated source such as JWT
		// Keeping it in the body for the sake of brevity
		br := cs.BookingRequest{}
		if err := c.ShouldBindJSON(&br); err != nil {
			c.JSON(http.StatusBadRequest, newErrResp(invalidRequestBody))
			return
		}
		id, err := b.Book(br)
		if userErr, ok := err.(cs.UserErr); ok {
			c.JSON(http.StatusBadRequest, newErrResp(userErr.Error()))
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, newErrResp(serverErrMsg))
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
