package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	cs "github.com/developersam1995/cabs-service/lib/cab-service"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func TestBooking(t *testing.T) {
	db := cs.NewBookingsMockDb()

	r := gin.Default()
	bs := cs.NewBookingService(db)

	setupBookingRoutes(r, bs)

	pw := httptest.NewRecorder()

	reqBody, err := json.Marshal(jsonMock{
		"from_lat":    12.44,
		"from_lon":    66.23,
		"to_lat":      13.232,
		"to_lon":      33.34,
		"user_id":     2,
		"pickup_time": 1524242631,
	})
	if err != nil {
		t.Fatal("Invalid body in test")
	}
	req, _ := http.NewRequest("POST", "/booking", bytes.NewBuffer(reqBody))
	r.ServeHTTP(pw, req)
	assert.Equal(t, http.StatusCreated, pw.Code, "Booking not created")

	gw := httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/booking/1", nil)
	r.ServeHTTP(gw, req)

	assert.Equal(t, http.StatusOK, gw.Code, "Unable to fetch bookings")

	// create one more booking
	newReq, _ := http.NewRequest("POST", "/booking", bytes.NewBuffer(reqBody))
	newRes := httptest.NewRecorder()
	r.ServeHTTP(newRes, newReq)
	assert.Equal(t, http.StatusBadRequest, newRes.Code, "Booking created twice")
}
