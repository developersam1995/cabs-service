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
		if config.Pooling {
			db.d.SetMaxOpenConns(config.Connections)
		} else {
			db.d.SetMaxOpenConns(1)
		}
		db.d.MustExec(usersSchema)
		db.d.MustExec(cabsSchema)
		db.d.MustExec(bookingsSchema)
	})

	return db, err

}

func (db *Db) SaveBooking(req cs.BookingRequest) (int, error) {
	r, err := db.d.Exec(`INSERT INTO bookings (from_lat, from_lon, to_lat, to_lon, pickup_time, user_id)
	VALUES (?, ?, ?, ?, ?, ?)`, req.FromLat, req.FromLon, req.ToLat, req.ToLon, req.PickupTime, req.UserID)
	id, _ := r.LastInsertId()
	return int(id), err
}

func (db *Db) FetchBookings(userID int) ([]cs.BookingRequest, error) {
	brs := []cs.BookingRequest{}
	err := db.d.Select(&brs, `SELECT id, from_lat, from_lon,
	to_lat, to_lon, pickup_time, is_confirmed, user_id 
	FROM bookings ORDER BY id DESC`)
	return brs, err
}

func (db *Db) FetchCabs(lat, lon float64, distance int) ([]cs.Cab, error) {
	cabs := []cs.Cab{}
	// calc distance and filter
	err := db.d.Select(&cabs, `SELECT veh_no, lat, lon FROM cabs 
						HAVING (6371 * acos(cos(radians(?)) * cos(radians(lat)) *
			            cos(radians(lon) - radians(?)) + sin(radians(?)) * sin(radians(lat)))) < ?`, lat, lon, lat, distance)
	return cabs, err
}

func (db *Db) FetchUnconfirmedBookings(userID int) ([]cs.BookingRequest, error) {
	brs := []cs.BookingRequest{}
	err := db.d.Select(&brs, `SELECT id, from_lat, from_lon,
	to_lat, to_lon, pickup_time, is_confirmed, user_id 
	FROM bookings WHERE is_confirmed=0 AND user_id = ? ORDER BY id DESC`, userID)
	return brs, err
}
