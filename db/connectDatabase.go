package db

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "sql12741017:rMjHllcWYj@tcp(sql12.freesqldatabase.com:3306)/sql12741017")

	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_test")
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return db, nil
}
