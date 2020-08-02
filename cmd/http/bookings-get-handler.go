package main

import (
	"net/http"
	"strconv"

	cs "github.com/developersam1995/cabs-service/lib/cab-service"
	"github.com/gin-gonic/gin"
)

func MakeBookingsGetter(b cs.BookingService) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Should use userID from an authenticated source such as JWT
		// Keeping it in the body for the sake of brevity
		userID := c.Param("userID")
		uID, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": invalidParams,
			})
			return
		}
		brs, err := b.ListAll(uID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": serverErrMsg,
			})
			return
		}
		c.JSON(
			http.StatusOK,
			gin.H{
				"bookings": brs,
			},
		)
	}
}
