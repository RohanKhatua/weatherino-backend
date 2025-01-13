package admin

import (
	"go-weather/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewAllDataHandler(c *gin.Context) {

	// Retrieve and validate query parameters
	measurement := c.Query("measurement")
	if measurement == "" {
		utils.SendError(http.StatusBadRequest, "Measurement value is required", c)
		return
	}

	// Show all records under the measurement
	err := utils.ShowAllRecordsUnderMeasurement(measurement)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, "Query from InfluxDB failed", c)
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message": "Data queried successfully | Check logs for details",
	})
}
