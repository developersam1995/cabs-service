package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"

	"github.com/developersam1995/cabs-service/storage"
)

type jsonMock map[string]interface{}

func TestBooking(t *testing.T) {
	db := storage.New(storage.DbConfig{})
	router := setupRouter(db)

	w := httptest.NewRecorder()
	reqBody, err := json.Marshal(jsonMock{
		"from": jsonMock{
			"lat": 12.44,
			"lon": 66.23,
		},
		"to": jsonMock{
			"lat": 13.232,
			"lon": 33.34,
		},
		"user_id":     1123,
		"pickup_time": 15223623633,
	})
	if err != nil {
		log.Fatalln("Invalid body in test")
	}
	req, _ := http.NewRequest("POST", "/booking", bytes.NewBuffer(reqBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Booking not created")

}
