package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	cs "github.com/developersam1995/cabs-service/lib/cab-service"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func TestGetCabs(t *testing.T) {
	newCabs := []cs.Cab{
		cs.Cab{
			VehNo: "123",
			Lat:   12.11,
			Lon:   66.12,
		},
		cs.Cab{
			VehNo: "124",
			Lat:   12.11,
			Lon:   66.12,
		},
		cs.Cab{
			VehNo: "125",
			Lat:   12.11,
			Lon:   66.12,
		},
	}

	db := cs.NewCabsMockDb(newCabs)

	r := gin.Default()
	cs := cs.NewCabsService(db)

	setupCabRoutes(r, cs)

	gw := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/cabs?lat=12.11&lon=66.12&dist=500", nil)
	r.ServeHTTP(gw, req)
	assert.Equal(t, http.StatusOK, gw.Code, "Could not list cabs")

}
