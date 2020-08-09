package cs

import (
	"errors"
)

// These are errors to show to the user
type UserErr error

var ErrUnconfirmedBookingsExist UserErr = errors.New("There are unconfirmed bookings. Please wait while we confirm")
var ErrNoSuchUser UserErr = errors.New("There is no user with such userID")
