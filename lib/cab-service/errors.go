package cs

import (
	"errors"
)

var ErrUnconfirmedBookingsExist = errors.New("There are unconfirmed bookings. Please wait while we confirm")
var ErrNoSuchUser = errors.New("There is no user with such userId")
