package cs

type Location struct {
	Lat float64
	Lon float64
}

type BookRequest struct {
	From       Location
	To         Location
	UserID     int
	PickupTime int
}

// Some storage interface needs to implement this
type bookingSaver interface {
	SaveBooking(BookRequest) (id int, err error)
}

// Services provided by the package
type BookingService interface {
	Book(BookRequest) (id int, err error)
}

type store struct {
	db bookingSaver
}

func NewBookingService(repo bookingSaver) BookingService {
	return &store{
		db: repo,
	}

}

func (s *store) Book(r BookRequest) (int, error) {
	// Any business logic like validation goes here
	return s.db.SaveBooking(r)
}
