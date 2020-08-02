package storage

import (
	"sync"

	cs "github.com/developersam1995/cabs-service/lib/cab-service"
)

var db *Db
var once sync.Once

type Db struct {
}

// Configuration needed for the initialization of connections to db
type DbConfig struct {
}

func New(config DbConfig) *Db {
	once.Do(func() {
		db = &Db{}
	})
	return db
}

func (db *Db) SaveBooking(req cs.BookRequest) (int, error) {
	return 1, nil
}
