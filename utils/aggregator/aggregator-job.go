package aggregator

import (
	customlogger "go-weather/custom-logger"
	dbutils "go-weather/utils/db-utils"
	"time"
)

func DailyAggregation() {
	now := time.Now()
	prevDayStart := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.Local)
	prevDayEnd := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	// TODO: Ensure no parallel execution

	// Get the average from sensor_data
	singleDayWeatherAverage, err := GetWeatherAveragesByTimeRange(prevDayStart, prevDayEnd, "sensor_data")
	if err != nil {
		customlogger.Logger.Errorf("Error getting weather averages: %s", err)
		return
	}

	// Write aggregated data
	err = dbutils.WritePoint(float32(singleDayWeatherAverage.Temperature), float32(singleDayWeatherAverage.Humidity), "daily_averages", prevDayStart)
	if err != nil {
		customlogger.Logger.Errorf("Error writing daily average metrics: %s", err)
		return
	}
	customlogger.Logger.Infof("Aggregation Point Written - Temperature: %f Humidity: %f", singleDayWeatherAverage.Temperature, singleDayWeatherAverage.Humidity)

	// Delete previous day's data
	err = dbutils.DeleteDailyData("sensor_data")
	if err != nil {
		customlogger.Logger.Errorf("Error deleting daily data: %s", err)
		return
	}
	customlogger.Logger.Info("Daily data successfully deleted.")
}
