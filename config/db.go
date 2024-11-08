package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "testing_go:hAmA^7]eZ#B}@tcp(68.178.150.12:3306)/rngdev_testing_go")
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_test")

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Connected to MySQL")

	return db, nil
}
