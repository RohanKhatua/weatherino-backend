package utils

import (
	"fmt"
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
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("INFLUXDB_ADMIN_TOKEN")
	url := fmt.Sprintf("http://localhost:%s", os.Getenv("INFLUXDB_PORT"))

	Client = influxdb2.NewClient(url, token)
	WriteAPI = Client.WriteAPIBlocking(os.Getenv("INFLUXDB_INIT_ORG"), os.Getenv("INFLUXDB_INIT_BUCKET"))
	QueryAPI = Client.QueryAPI(os.Getenv("INFLUXDB_INIT_ORG"))
	DeleteAPI = Client.DeleteAPI()
}
