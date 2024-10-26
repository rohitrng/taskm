package routes

import (
	"rngdev/handler"

	"github.com/gin-gonic/gin"
)

func api() {
	r := gin.Default()
	r.GET("/", handler.Login)
	r.Run(":8080")
}