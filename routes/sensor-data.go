package routes

import (
	customlogger "go-weather/custom-logger"
	"go-weather/models"
	"go-weather/utils/aggregator"
	dbutils "go-weather/utils/db-utils"
	serverutils "go-weather/utils/server-utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SaveSensorDataHandler handles the writing of sensor data to InfluxDB
func SaveSensorDataHandler(c *gin.Context) {
	// Set request time as soon as this function gets invoked for accuracy

	requestTime := time.Now()

	// Retrieve and validate query parameters
	temperature := c.Query("temp")
	humidity := c.Query("humid")

	if temperature == "" || humidity == "" {
		serverutils.SendError(http.StatusBadRequest, "Temperature and Humidity values are required", c)
		return
	}

	// Parse temperature
	temperatureValue, err := strconv.ParseFloat(temperature, 32)
	if err != nil {
		serverutils.SendError(http.StatusBadRequest, "Invalid Temperature value", c)
		return
	}

	// Parse humidity
	humidityValue, err := strconv.ParseFloat(humidity, 32)
	if err != nil {
		serverutils.SendError(http.StatusBadRequest, "Invalid Humidity value", c)
		return
	}

	// Write the data to InfluxDB
	err = dbutils.WritePoint(float32(temperatureValue), float32(humidityValue), "sensor_data", requestTime)
	if err != nil {
		serverutils.SendError(http.StatusInternalServerError, "Write to InfluxDB failed", c)
		return
	}

	customlogger.Logger.Infof("Sensor || Temperature: %f, Humidity: %f", temperatureValue, humidityValue)

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message": "Data written successfully",
	})
}

func RetrieveDataByTimeDuration(c *gin.Context) {
	hours := c.Query("hours")
	minutes := c.Query("minutes")
	seconds := c.Query("seconds")

	queryDuration := models.WeatherDuration{}

	if hours != "" {
		hoursInt, err := strconv.Atoi(hours)
		if err != nil {
			serverutils.SendError(http.StatusBadRequest, "Invalid hours value", c)
			return
		}
		queryDuration.Hours = hoursInt
	}

	if minutes != "" {
		minutesInt, err := strconv.Atoi(minutes)
		if err != nil {
			serverutils.SendError(http.StatusBadRequest, "Invalid minutes value", c)
			return
		}
		queryDuration.Minutes = minutesInt
	}

	if seconds != "" {
		secondsInt, err := strconv.Atoi(seconds)
		if err != nil {
			serverutils.SendError(http.StatusBadRequest, "Invalid seconds value", c)
			return
		}
		queryDuration.Seconds = secondsInt
	}

	// customlogger.Logger.Println("Query Duration: ", queryDuration)
	averages, err := aggregator.GetWeatherAveragesByDuration(queryDuration, "sensor_data")

	if err != nil {
		serverutils.SendError(http.StatusInternalServerError, "Failed to retrieve data", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"temperature": averages.Temperature,
		"humidity":    averages.Humidity,
	})
}
