package main

import (
	cl "go-weather/custom-logger"
	"go-weather/routes"
	"go-weather/routes/admin"
	"go-weather/utils"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func setupRouter() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	// this disables logging - use custom logger for clarity and this provides performance benefits
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

func setupCron() {
	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", utils.DailyAggregation)
	// Setting Cron Job to run every minute for testing
	// Use postman to send requests for ten minutes and check

	if err != nil {
		cl.Logger.Errorf("Error adding cron job: %s", err)
		return
	}

	c.Start()
}

func main() {
	// Initialize the logger
	cl.Init()

	// Create the InfluxDB client and initialize the API objects
	utils.CreateInfluxClient()
	setupCron()
	setupRouter()

}
