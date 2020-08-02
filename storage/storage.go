package storage

import (
	"bytes"
	"sync"

	cs "github.com/developersam1995/cabs-service/lib/cab-service"
	"github.com/jmoiron/sqlx"
)

var db *Db
var once sync.Once

type Db struct {
	d *sqlx.DB
}

// Configuration needed for the initialization of connections to db
type DbConfig struct {
	Host        string `json:"host"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Db          string `json:"db"`
	Pooling     bool   `json:"pooling"`
	Connections int    `json:"connections"`
}

func New(config DbConfig) (*Db, error) {
	var err error
	once.Do(func() {
		var dsn bytes.Buffer
		dsn.WriteString(config.User + ":" + config.Password + "@tcp(" + config.Host +
			":3306)/" + config.Db)
		d, er := sqlx.Connect("mysql", dsn.String())
		err = er
		db = &Db{
			d: d,
		}
	})

	return db, err

}

func (db *Db) SaveBooking(req cs.BookingRequest) (int, error) {
	return 1, nil
}

func (db *Db) FetchBookings(userID int) ([]cs.BookingRequest, error) {
	brs := []cs.BookingRequest{}
	return brs, nil
}

func (db *Db) FetchCabs(l cs.Location, distance int) ([]cs.Cabs, error) {
	cabs := []cs.Cabs{}
	return cabs, nil
}
