package admin

import (
	"go-weather/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewAllDataHandler(c *gin.Context) {
	// Helper function for returning errors
	sendError := func(status int, message string) {
		c.JSON(status, gin.H{"error": message})
	}

	// Retrieve and validate query parameters
	measurement := c.Query("measurement")
	if measurement == "" {
		sendError(http.StatusBadRequest, "Measurement value is required")
		return
	}

	// Show all records under the measurement
	err := utils.ShowAllRecordsUnderMeasurement(measurement)
	if err != nil {
		sendError(http.StatusInternalServerError, "Query from InfluxDB failed")
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message": "Data queried successfully | Check logs for details",
	})
}
