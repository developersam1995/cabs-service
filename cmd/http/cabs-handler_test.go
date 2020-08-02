package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	cs "github.com/developersam1995/cabs-service/lib/cab-service"
	"github.com/developersam1995/cabs-service/storage"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func TestGetCabs(t *testing.T) {
	db, err := storage.New(config.Db)

	if err != nil {
		t.Fatal(err)
	}

	r := gin.Default()
	cs := cs.NewCabsService(db)

	setupCabRoutes(r, cs)

	gw := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/cabs?lat=12.23&lon=66.121&dist=500", nil)
	r.ServeHTTP(gw, req)
	assert.Equal(t, http.StatusOK, gw.Code, "Could not list cabs")

}
