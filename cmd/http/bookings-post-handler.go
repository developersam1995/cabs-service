package main

import (
	"net/http"

	cs "github.com/developersam1995/cabs-service/lib/cab-service"
	"github.com/gin-gonic/gin"
)

func makeBookingsPoster(b cs.BookingService) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Should use userID from an authenticated source such as JWT
		// Keeping it in the body for the sake of brevity
		br := cs.BookingRequest{}
		if err := c.ShouldBindJSON(&br); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := b.Book(br)
		if err == cs.ErrUnconfirmedBookingsExist {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": serverErrMsg,
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
