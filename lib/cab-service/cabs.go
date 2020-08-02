package cs

type Cabs struct {
	VehNo       string   `json:"veh_no" binding:"required"`
	Position    Location `json:"from" binding:"required"`
	DriverName  string   `json:"driver_name" binding:"required"`
	DriverPhone int      `json:"driver_phone" binding:"required"`
}

// Some storage interface needs to implement this
type cabsRepo interface {
	FetchCabs(l Location, distance int) ([]Cabs, error)
}

// Services provided by the package
type CabsService interface {
	ListAll(l Location, distance int) ([]Cabs, error)
}

type cabsStore struct {
	db cabsRepo
}

func NewCabsService(repo cabsRepo) *cabsStore {
	return &cabsStore{
		db: repo,
	}

}

func (s *cabsStore) ListAll(l Location, d int) ([]Cabs, error) {
	// Any business logic like validation goes here
	return s.db.FetchCabs(l, d)
}
