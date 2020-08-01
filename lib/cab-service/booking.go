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
