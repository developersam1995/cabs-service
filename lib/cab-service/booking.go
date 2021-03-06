package cs

type BookingRequest struct {
	ID          int     `json:"id" db:"id"`
	FromLat     float64 `json:"from_lat" binding:"required" db:"from_lat"`
	FromLon     float64 `json:"from_lon" binding:"required" db:"from_lon"`
	ToLat       float64 `json:"to_lat" binding:"required" db:"to_lat"`
	ToLon       float64 `json:"to_lon" binding:"required" db:"to_lon"`
	UserID      int     `json:"user_id" binding:"required" db:"user_id"`
	PickupTime  int     `json:"pickup_time" binding:"required" db:"pickup_time"`
	IsConfirmed int     `json:"is_confirmed" db:"is_confirmed"`
}

// Some storage interface needs to implement this
type bookingRepo interface {
	SaveBooking(BookingRequest) (id int, err error)
	FetchBookings(userID int) ([]BookingRequest, error)
	FetchUnconfirmedBookings(userID int) ([]BookingRequest, error)
}

// Services provided by the package
type BookingService interface {
	Book(BookingRequest) (id int, err error)
	ListAll(userID int) ([]BookingRequest, error)
}

type bookingStore struct {
	db bookingRepo
}

func NewBookingService(repo bookingRepo) BookingService {
	return &bookingStore{
		db: repo,
	}

}

func (s *bookingStore) Book(r BookingRequest) (int, error) {
	// Any business logic like validation goes here
	existingRequests, err := s.db.FetchUnconfirmedBookings(r.UserID)
	if err != nil {
		return 0, err
	}
	if len(existingRequests) > 0 {
		return 0, ErrUnconfirmedBookingsExist
	}

	return s.db.SaveBooking(r)
}

func (s *bookingStore) ListAll(userID int) ([]BookingRequest, error) {
	// Any business logic goes here
	return s.db.FetchBookings(userID)
}
