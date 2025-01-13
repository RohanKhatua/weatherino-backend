package routes

import (
	customlogger "go-weather/custom-logger"
	"go-weather/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SaveSensorDataHandler handles the writing of sensor data to InfluxDB
func SaveSensorDataHandler(c *gin.Context) {
	// Retrieve and validate query parameters
	temperature := c.Query("temp")
	humidity := c.Query("humid")

	if temperature == "" || humidity == "" {
		utils.SendError(http.StatusBadRequest, "Temperature and Humidity values are required", c)
		return
	}

	// Parse temperature
	temperatureValue, err := strconv.ParseFloat(temperature, 32)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "Invalid Temperature value", c)
		return
	}

	// Parse humidity
	humidityValue, err := strconv.ParseFloat(humidity, 32)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "Invalid Humidity value", c)
		return
	}

	// Write the data to InfluxDB
	err = utils.WritePoint(float32(temperatureValue), float32(humidityValue))
	if err != nil {
		utils.SendError(http.StatusInternalServerError, "Write to InfluxDB failed", c)
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message": "Data written successfully",
	})
}

func RetrieveDataByTimeDuration(c *gin.Context) {
	hours := c.Query("hours")
	minutes := c.Query("minutes")
	seconds := c.Query("seconds")

	queryDuration := utils.WeatherDuration{}

	if hours != "" {
		hoursInt, err := strconv.Atoi(hours)
		if err != nil {
			utils.SendError(http.StatusBadRequest, "Invalid hours value", c)
			return
		}
		queryDuration.Hours = hoursInt
	}

	if minutes != "" {
		minutesInt, err := strconv.Atoi(minutes)
		if err != nil {
			utils.SendError(http.StatusBadRequest, "Invalid minutes value", c)
			return
		}
		queryDuration.Minutes = minutesInt
	}

	if seconds != "" {
		secondsInt, err := strconv.Atoi(seconds)
		if err != nil {
			utils.SendError(http.StatusBadRequest, "Invalid seconds value", c)
			return
		}
		queryDuration.Seconds = secondsInt
	}

	customlogger.Logger.Println("Query Duration: ", queryDuration)
	averages, err := utils.GetWeatherAveragesByDuration(queryDuration)

	if err != nil {
		utils.SendError(http.StatusInternalServerError, "Failed to retrieve data", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"temperature": averages.Temperature,
		"humidity":    averages.Humidity,
	})
}
