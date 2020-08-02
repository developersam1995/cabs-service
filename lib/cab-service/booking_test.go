package cs

import (
	"testing"
)

func TestBookingTwice(t *testing.T) {
	mockDb := NewBookingsMockDb()
	serv := NewBookingService(mockDb)

	serv.Book(BookingRequest{
		UserID: 1,
	})

	_, err := serv.Book(BookingRequest{
		UserID: 1,
	})

	if err != ErrUnconfirmedBookingsExist {
		t.Error("Same user can book twice")
	}
}
