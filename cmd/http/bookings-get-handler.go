package main

import (
	"net/http"
	"strconv"

	cs "github.com/developersam1995/cabs-service/lib/cab-service"
	"github.com/gin-gonic/gin"
)

type Booking struct {
	ID          int     `json:"id"`
	FromLat     float64 `json:"from_lat" binding:"required"`
	FromLon     float64 `json:"from_lon" binding:"required"`
	ToLat       float64 `json:"to_lat" binding:"required"`
	ToLon       float64 `json:"to_lon" binding:"required"`
	UserID      int     `json:"user_id" binding:"required"`
	PickupTime  int     `json:"pickup_time" binding:"required"`
	IsConfirmed int     `json:"is_confirmed"`
}

// @Summary Get the list of cab bookings previously made by the user
// @Produce  json
// @Param userID path integer true "userID"
// @Success 200 {object} []Booking
// @Failure 400 {object} ErrResp
// @Router /booking/{userID} [get]
func makeBookingsGetter(b cs.BookingService) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Should use userID from an authenticated source such as JWT
		// Keeping it in the route for the sake of brevity
		userID := c.Param("userID")
		uID, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, newErrResp(invalidParams))
			return
		}
		brs, err := b.ListAll(uID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, newErrResp(serverErrMsg))
			return
		}
		c.JSON(
			http.StatusOK,
			brs,
		)
	}
}
