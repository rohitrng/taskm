package handler

import (
	"net/http"
	"rngdev/db"

	"github.com/gin-gonic/gin"
)

type Emp struct {
	Username string `json:"name"`
	Age      string `json:"age"`
}

func Login(c *gin.Context) {
	db, err := db.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	res, err := db.Query("select username , age from users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	var emps []Emp
	for res.Next() {
		var emp Emp
		emps = append(emps, Emp{emp.Username, emp.Age})
	}
	c.JSON(http.StatusOK, gin.H{"message": emps})
}
