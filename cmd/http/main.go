package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	cs "github.com/developersam1995/cabs-service/lib/cab-service"
	"github.com/developersam1995/cabs-service/storage"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/developersam1995/cabs-service/cmd/http/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type jsonMock map[string]interface{}

var config configuration

func setupBookingRoutes(r *gin.Engine, b cs.BookingService) {
	r.POST("/booking", makeBookingsPoster(b))
	r.GET("/booking/:userID", makeBookingsGetter(b))
}

func setupCabRoutes(r *gin.Engine, c cs.CabsService) {
	r.GET("/cabs", makeCabsGetter(c))
}

func init() {

	config = configuration{
		Port: 8080,
		Db: storage.DbConfig{
			Host:     "127.0.0.1",
			User:     "server_x",
			Password: "testpassword",
			Db:       "cabs_db",
		},
		MaxCabGetDist: 2000,
	}

	if len(os.Args) > 1 {
		d, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			log.Println("Invalid argument for config file")
		}
		err = json.Unmarshal(d, &config)
		if err != nil {
			log.Println("Invalid json format. Please refer config_example.json")
		}
	}

}

// @title Cabs Servies API
// @version 1.0
// @description APIs for booking and getting cabs.

// @host localhost:8080

func main() {
	r := gin.Default()
	db, err := storage.New(config.Db)
	if err != nil {
		log.Fatalln("Cannot connect to db:", err)
	}

	bookingService := cs.NewBookingService(db)
	cabsService := cs.NewCabsService(db)

	setupBookingRoutes(r, bookingService)
	setupCabRoutes(r, cabsService)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	p := ":8080"
	if config.Port > 0 {
		p = ":" + strconv.Itoa(config.Port)
	}
	r.Run(p)
}
