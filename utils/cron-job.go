package utils

import (
	customlogger "go-weather/custom-logger"
	"time"
)

func DailyAggregation() {
	// setting this to 1 minute to test functionality
	singleDayDuration := WeatherDuration{
		Hours: 24,
	}
	var singleDayWeatherAverage WeatherAverages

	// get the average from sensor_data and then write it to daily averages
	singleDayWeatherAverage, err := GetWeatherAveragesByDuration(singleDayDuration, "sensor_data")

	if err != nil {
		customlogger.Logger.Errorf("Error getting weather averages: %s", err)
		return
	}

	// We subtract a day as the cron job runs at 00:00 of the next day
	err = WritePoint(float32(singleDayWeatherAverage.Temperature), float32(singleDayWeatherAverage.Humidity), "daily_averages", time.Now().AddDate(0, 0, -1))

	if err != nil {
		customlogger.Logger.Errorf("Error writing daily average metrics: %s", err)
	}

	customlogger.Logger.Infof("Aggregation Point Written - Temperature: %f Humidity: %f", singleDayWeatherAverage.Temperature, singleDayWeatherAverage.Humidity)
}
