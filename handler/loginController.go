package handler

import (
	"net/http"
	"rngdev/db"

	"github.com/gin-gonic/gin"
)

type Emp struct {
	Username string `json:"name"`
	Id       string `json:"id"`
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	db, err := db.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	res, err := db.Query("select username, id from users where username = ? and password = ?", username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	var emp Emp
	if res.Next() {
		err := res.Scan(&emp.Username, &emp.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": emp})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
	}
}
