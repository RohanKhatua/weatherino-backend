package main

import (
	cl "go-weather/custom-logger"
	"go-weather/routes"
	"go-weather/routes/admin"
	"go-weather/utils"

	"github.com/gin-gonic/gin"
)

func setupRouter() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Admin routes
	r.GET("/admin/view-data", admin.ViewAllDataHandler)
	r.DELETE("/admin/delete-data", admin.DeleteDataHandler)

	// Sensor data route
	r.POST("/sensor-data", routes.SaveSensorDataHandler)
	r.GET("/sensor-data", routes.RetrieveDataByTimeDuration)

	// Run the server
	err := r.Run(":8080")
	if err != nil {
		cl.Logger.Fatal(err)
	}

}

func main() {
	// Initialize the logger
	cl.Init()

	// Create the InfluxDB client and initialize the API objects
	utils.CreateInfluxClient()
	setupRouter()

}
