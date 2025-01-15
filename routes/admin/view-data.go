package admin

import (
	dbutils "go-weather/utils/db-utils"
	serverutils "go-weather/utils/server-utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewAllDataHandler(c *gin.Context) {

	// Retrieve and validate query parameters
	measurement := c.Query("measurement")
	if measurement == "" {
		serverutils.SendError(http.StatusBadRequest, "Measurement value is required", c)
		return
	}

	// Show all records under the measurement
	err := dbutils.ShowAllRecordsUnderMeasurement(measurement)
	if err != nil {
		serverutils.SendError(http.StatusInternalServerError, "Query from InfluxDB failed", c)
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message": "Data queried successfully | Check logs for details",
	})
}
