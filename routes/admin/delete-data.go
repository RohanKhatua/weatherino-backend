package admin

import (
	"go-weather/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteDataHandler(c *gin.Context) {
	// Helper function for returning errors
	sendError := func(status int, message string) {
		c.JSON(status, gin.H{"error": message})
	}

	// Delete all data
	err := utils.DeleteAllData()
	if err != nil {
		sendError(http.StatusInternalServerError, "Delete from InfluxDB failed")
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message": "Data deleted successfully",
	})
}
