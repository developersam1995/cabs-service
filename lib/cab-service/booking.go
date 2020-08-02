package cs

type Location struct {
	Lat float64 `json:"lat" binding:"required"`
	Lon float64 `json:"lon" binding:"required"`
}

type BookingRequest struct {
	From       Location `json:"from" binding:"required"`
	To         Location `json:"to" binding:"required"`
	UserID     int      `json:"user_id" binding:"required"`
	PickupTime int      `json:"pickup_time" binding:"required"`
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

type store struct {
	db bookingRepo
}

func NewBookingService(repo bookingRepo) BookingService {
	return &store{
		db: repo,
	}

}

func (s *store) Book(r BookingRequest) (int, error) {
	// Any business logic like validation goes here
	return s.db.SaveBooking(r)
}

func (s *store) ListAll(userID int) ([]BookingRequest, error) {
	// Any business logic goes here
	return s.db.FetchBookings(userID)
}
