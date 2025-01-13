package utils

import (
	"context"
	"fmt"
	customlogger "go-weather/custom-logger"
	"log"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func WritePoint(temperature float32, humidity float32) error {
	p := influxdb2.NewPoint(
		"weather-data",
		map[string]string{"sensor": "DHT11"},
		map[string]interface{}{"temperature": temperature, "humidity": humidity},
		time.Now())
	err := WriteAPI.WritePoint(context.Background(), p)

	if err != nil {
		log.Printf("Write error: %s\n", err)
		return err
	}

	return nil
}

func DeleteAllData() error {
	org := os.Getenv("INFLUXDB_INIT_ORG")
	bucket := os.Getenv("INFLUXDB_INIT_BUCKET")
	start := time.Unix(0, 0)
	stop := time.Now()
	err := DeleteAPI.DeleteWithName(context.Background(), org, bucket, start, stop, "")

	if err != nil {
		log.Printf("Delete error: %s\n", err)
		return err
	}

	return nil
}

func ShowAllRecordsUnderMeasurement(mesaurement string) error {
	bucket := os.Getenv("INFLUXDB_INIT_BUCKET")
	query := fmt.Sprintf(
		`from(bucket:"%s") 
		|> range(start: 0) 
		|> filter(fn: (r) => r._measurement == "%s")`,
		bucket,
		mesaurement)

	customlogger.Logger.Info(query)

	results, err := QueryAPI.Query(context.Background(), query)
	if err != nil {
		log.Printf("Query error: %s\n", err)
		return err
	}

	for results.Next() {
		fieldName := results.Record().Field()
		fieldValue := results.Record().Value()
		resultTime := results.Record().Time()
		customlogger.Logger.Info(fmt.Sprintf("Field: %s, Value: %v at %s", fieldName, fieldValue, resultTime.String()))
	}

	return nil
}
