package dbutils

import (
	"context"
	"fmt"
	customlogger "go-weather/custom-logger"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func WritePoint(temperature float32, humidity float32, measurement string, timestamp time.Time) error {
	p := influxdb2.NewPoint(
		measurement,         // Measurement Name (Table Name)
		map[string]string{}, // Tags - not required
		map[string]interface{}{"temperature": temperature, "humidity": humidity},
		timestamp)
	err := WriteAPI.WritePoint(context.Background(), p)

	if err != nil {
		customlogger.Logger.Errorf("Write Error %s\n", err)
		return err
	}

	return nil
}

// DeleteAllData deletes all the data in the entire bucket - all measurements inside it as well.
func DeleteAllData() error {
	org := os.Getenv("INFLUXDB_INIT_ORG")
	bucket := os.Getenv("INFLUXDB_INIT_BUCKET")
	start := time.Unix(0, 0)
	stop := time.Now()
	err := DeleteAPI.DeleteWithName(context.Background(), org, bucket, start, stop, "")

	if err != nil {
		customlogger.Logger.Errorf("Delete Failed %s\n", err)
		return err
	}

	return nil
}

func ShowAllRecordsUnderMeasurement(measurement string) error {
	bucket := os.Getenv("INFLUXDB_INIT_BUCKET")
	query := fmt.Sprintf(
		`from(bucket:"%s") 
		|> range(start: 0) 
		|> filter(fn: (r) => r._measurement == "%s")`,
		bucket,
		measurement)

	customlogger.Logger.Infof("Query : %s", query)

	results, err := QueryAPI.Query(context.Background(), query)
	if err != nil {
		customlogger.Logger.Errorf("Query Error %s\n", err)
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

func DeleteDailyData(measurement string) error {
	bucket := os.Getenv("INFLUXDB_INIT_BUCKET")
	org := os.Getenv("INFLUXDB_INIT_ORG")

	if bucket == "" || org == "" {
		return fmt.Errorf("INFLUXDB_INIT_BUCKET or INFLUXDB_INIT_ORG environment variable is not set")
	}

	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, now.Location())
	stop := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	err := DeleteAPI.DeleteWithName(context.Background(), org, bucket, start, stop, measurement)
	if err != nil {
		customlogger.Logger.Errorf("Failed to delete daily data for measurement '%s': %s", measurement, err)
		return err
	}

	customlogger.Logger.Infof("Successfully deleted daily data for measurement '%s' from %s to %s", measurement, start, stop)
	return nil
}
