package main

import (
	"net/http"
	"strconv"

	cs "github.com/developersam1995/cabs-service/lib/cab-service"
	"github.com/gin-gonic/gin"
)

func MakeCabsGetter(cas cs.CabsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		lat := c.Query("lat")
		lon := c.Query("lon")
		dist := c.Query("dist")

		latitude, latErr := strconv.ParseFloat(lat, 64)
		longitude, lonErr := strconv.ParseFloat(lon, 64)
		distance, distErr := strconv.Atoi(dist)

		if latErr != nil || lonErr != nil || distErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": invalidQuery,
			})
			return
		}

		loc := cs.Location{
			Lat: latitude,
			Lon: longitude,
		}
		cabs, err := cas.ListAll(loc, distance)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": serverErrMsg,
			})
			return
		}
		c.JSON(
			http.StatusOK,
			gin.H{
				"cabs": cabs,
			},
		)
	}
}
