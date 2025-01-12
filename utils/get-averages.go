package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

type WeatherDuration struct {
	Hours   int
	Minutes int
	Seconds int
}

type WeatherAverages struct {
	Temperature float64
	Humidity    float64
}

// GetAverages is a function that returns the average temperature and humidity for the last specified time period
// The function takes a WeatherDuration as an argument and returns a WeatherAverages struct
// if no duration is passed - the function defaults to 10 minutes
func GetWeatherAveragesByDuration(duration WeatherDuration) (WeatherAverages, error) {
	if duration.Hours == 0 && duration.Minutes == 0 && duration.Seconds == 0 {
		duration.Minutes = 10
	}

	totalDuration := time.Duration(duration.Hours)*time.Hour + time.Duration(duration.Minutes)*time.Minute + time.Duration(duration.Seconds)*time.Second

	// limit the total duration to 7 days as the DB has a retention policy of 7 days
	if totalDuration > 7*24*time.Hour {
		totalDuration = 7 * 24 * time.Hour
	}

	rangeStart := fmt.Sprintf("-%ds", int(totalDuration.Seconds()))
	bucket := os.Getenv("INFLUXDB_INIT_BUCKET")

	query := fmt.Sprintf(`from(bucket:"%s")
	|> range(start: %s)
	|> filter(fn: (r) => r._measurement == "weather-data")
	|> filter(fn: (r) => r._field == "temperature" or r._field == "humidity")
	|> mean()
	|> yield()`, bucket, rangeStart)

	results, err := QueryAPI.Query(context.Background(), query)
	if err != nil {
		log.Println(err)
		return WeatherAverages{}, err
	}

	averageTemperature := 0.0
	averageHumidity := 0.0

	for results.Next() {
		record := results.Record().Values()

		if record["_field"] == "temperature" {
			averageTemperature = record["_value"].(float64)
		}

		if record["_field"] == "humidity" {
			averageHumidity = record["_value"].(float64)
		}
	}

	return WeatherAverages{
		Temperature: averageTemperature,
		Humidity:    averageHumidity,
	}, nil
}
