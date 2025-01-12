package routes

import (
	"go-weather/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SaveSensorDataHandler handles the writing of sensor data to InfluxDB
func SaveSensorDataHandler(c *gin.Context) {
	// Helper function for returning errors
	sendError := func(status int, message string) {
		c.JSON(status, gin.H{"error": message})
	}

	// Retrieve and validate query parameters
	temperature := c.Query("temp")
	humidity := c.Query("humid")

	if temperature == "" || humidity == "" {
		sendError(http.StatusBadRequest, "Temperature and Humidity values are required")
		return
	}

	// Parse temperature
	temperatureValue, err := strconv.ParseFloat(temperature, 32)
	if err != nil {
		sendError(http.StatusBadRequest, "Invalid Temperature value")
		return
	}

	// Parse humidity
	humidityValue, err := strconv.ParseFloat(humidity, 32)
	if err != nil {
		sendError(http.StatusBadRequest, "Invalid Humidity value")
		return
	}

	// Write the data to InfluxDB
	err = utils.WritePoint(float32(temperatureValue), float32(humidityValue))
	if err != nil {
		sendError(http.StatusInternalServerError, "Write to InfluxDB failed")
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message": "Data written successfully",
	})
}
