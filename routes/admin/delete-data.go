package admin

import (
	customlogger "go-weather/custom-logger"
	"go-weather/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteDataHandler(c *gin.Context) {
	// Delete all data
	err := utils.DeleteAllData()
	if err != nil {
		utils.SendError(http.StatusInternalServerError, "Delete from InfluxDB failed", c)
		return
	}
	customlogger.Logger.Info("All Data Deleted")
	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message": "Data deleted successfully",
	})
}
