package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	router.POST("/message", func(c *gin.Context) {
		var msg Message

		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"status":  "success",
			"message": fmt.Sprintf("Received message from %s: %s", msg.Name, msg.Message),
		})
	})

	router.Run(":8080")
}
