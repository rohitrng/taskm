package controllers

import (
	"database/sql"
	"net/http"
	"rngdev/config"
	"rngdev/models"
	"rngdev/services"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	db, err := config.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	defer db.Close()
	err = services.Register(db, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Check(c *gin.Context) {
	db, err := sql.Open("mysql", "103.185.160.244:@tcp(localhost:3306)/task_management")
	// db, err := sql.Open("mysql", "rngdev_data:-~BGo-zuX5Yb@tcp(68.178.150.12:3306)/task_management")
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_test")

	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"message": err})
		return
	}

	if err = db.Ping(); err != nil {
		c.JSON(http.StatusCreated, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Connected to MySQL"})
}
