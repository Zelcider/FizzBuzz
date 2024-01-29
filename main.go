package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api/v1/ready", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	router.Run("localhost:8080")
}
