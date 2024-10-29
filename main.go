package main

import (
	"rngdev/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	routes.SetupRoutes(route)
	// route.Run(":8081")
	route.Run(":8080")

}
