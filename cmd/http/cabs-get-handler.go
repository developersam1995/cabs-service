package main

import (
	"net/http"
	"strconv"

	cs "github.com/developersam1995/cabs-service/lib/cab-service"
	"github.com/gin-gonic/gin"
)

type Cab struct {
	VehNo string  `json:"veh_no"`
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
}

// @Summary Get a list of cabs nearby from the given location
// @Produce  json
// @Param lat query number true "latitude"
// @Param lon query number true "longitude"
// @Param dist query integer true "distance in meters"
// @Failure 400 {object} ErrResp
// @Success 200 {object} []Cab
// @Router /cabs [get]
func makeCabsGetter(cas cs.CabsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		lat := c.Query("lat")
		lon := c.Query("lon")
		dist := c.Query("dist")

		latitude, latErr := strconv.ParseFloat(lat, 64)
		longitude, lonErr := strconv.ParseFloat(lon, 64)
		distance, distErr := strconv.Atoi(dist)

		if latErr != nil || lonErr != nil {
			c.JSON(http.StatusBadRequest, newErrResp(invalidQuery))
			return
		}

		if distErr != nil || distance > config.MaxCabGetDist {
			distance = config.MaxCabGetDist
		}

		cabs, err := cas.ListAll(latitude, longitude, distance)
		if err != nil {
			c.JSON(http.StatusInternalServerError, newErrResp(serverErrMsg))
			return
		}
		c.JSON(
			http.StatusOK,
			cabs,
		)
	}
}
