package utils

import (
	customlogger "go-weather/custom-logger"
	"log"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/joho/godotenv"
)

var Client influxdb2.Client
var WriteAPI api.WriteAPIBlocking
var QueryAPI api.QueryAPI
var DeleteAPI api.DeleteAPI

func CreateInfluxClient() {
	err := godotenv.Load()
	// we only handle the error in local environment
	// in production, docker compose will handle the error
	if os.Getenv("ENV") == "local" {
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	token := os.Getenv("INFLUXDB_ADMIN_TOKEN")
	url := os.Getenv("INFLUXDB_URL")

	Client = influxdb2.NewClient(url, token)
	WriteAPI = Client.WriteAPIBlocking(os.Getenv("INFLUXDB_INIT_ORG"), os.Getenv("INFLUXDB_INIT_BUCKET"))
	QueryAPI = Client.QueryAPI(os.Getenv("INFLUXDB_INIT_ORG"))
	DeleteAPI = Client.DeleteAPI()

	customlogger.Logger.Info("InfluxDB Client and APIs created")
}
