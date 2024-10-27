package routes

import (
	"rngdev/handler"

	"github.com/gin-gonic/gin"
)

func Api() {
	r := gin.Default()
	r.POST("/login", handler.Login)
	r.Run(":8080")
}
