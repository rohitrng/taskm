package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "rngdev_data:-~BGo-zuX5Yb@tcp(68.178.150.12:3306)/task_management")
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
