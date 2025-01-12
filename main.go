package main

import (
	"go-weather/routes"
	"go-weather/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create the InfluxDB client and initialize the API objects
	utils.CreateInfluxClient()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/sensor-data", routes.SaveSensorDataHandler)

	r.Run(":8080")

}
