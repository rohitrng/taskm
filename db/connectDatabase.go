package db

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "testing_go:Iq7%ylzK#3]D@tcp(68.178.150.12:3306)/rngdev_testing_go")
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return db, nil
}
