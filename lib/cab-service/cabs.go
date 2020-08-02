package cs

type Cabs struct {
	VehNo string  `json:"veh_no" binding:"required" db:"veh_no"`
	Lat   float64 `json:"lat" binding:"required" db:"lat"`
	Lon   float64 `json:"lon" binding:"required" db:"lon"`
}

// Some storage interface needs to implement this
type cabsRepo interface {
	FetchCabs(lat, lon float64, distance int) ([]Cabs, error)
}

// Services provided by the package
type CabsService interface {
	ListAll(lat, lon float64, distance int) ([]Cabs, error)
}

type cabsStore struct {
	db cabsRepo
}

func NewCabsService(repo cabsRepo) *cabsStore {
	return &cabsStore{
		db: repo,
	}

}

func (s *cabsStore) ListAll(lat, lon float64, d int) ([]Cabs, error) {
	// Any business logic like validation goes here
	return s.db.FetchCabs(lat, lon, d)
}
