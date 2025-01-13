package utils

import "github.com/gin-gonic/gin"

func SendError(status int, message string, c *gin.Context) {
	c.JSON(status, gin.H{
		"error": message,
	})
}
