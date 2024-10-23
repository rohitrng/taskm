package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Echo instance
	r := gin.Default()
	r.GET("/", Hello)
	r.Run(":8080")
}

// Handler
func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
