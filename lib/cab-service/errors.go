package cs

import (
	"errors"
)

var ErrUnconfirmedBookingsExist = errors.New("There are unconfirmed bookings. Please wait while we confirm")
