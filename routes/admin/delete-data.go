package admin

import (
	customlogger "go-weather/custom-logger"
	dbutils "go-weather/utils/db-utils"
	serverutils "go-weather/utils/server-utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteDataHandler(c *gin.Context) {
	// Delete all data
	err := dbutils.DeleteAllData()
	if err != nil {
		serverutils.SendError(http.StatusInternalServerError, "Bucket Data Delete Failed", c)
		return
	}
	customlogger.Logger.Info("All Data Deleted")
	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message": "Data deleted successfully",
	})
}
