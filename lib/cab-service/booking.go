package cs

type BookingRequest struct {
	FromLat    float64 `json:"from_lat" binding:"required"`
	FromLon    float64 `json:"from_lon: binding:required""`
	ToLat      float64 `json:"to_lat" binding:"required"`
	ToLon      float64 `json:"to_lon: binding:required""`
	UserID     int     `json:"user_id" binding:"required"`
	PickupTime int     `json:"pickup_time" binding:"required"`
}

// Some storage interface needs to implement this
type bookingRepo interface {
	SaveBooking(BookingRequest) (id int, err error)
	FetchBookings(userID int) ([]BookingRequest, error)
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
	return s.db.SaveBooking(r)
}

func (s *bookingStore) ListAll(userID int) ([]BookingRequest, error) {
	// Any business logic goes here
	return s.db.FetchBookings(userID)
}
