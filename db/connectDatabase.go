package db

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "rediss:AU6KAAIjcDFmM2JkODBkNmY4ZmU0YTI3OGQwYWFlMTY1NGRiNGE4M3AxMA@tcp(guided-guppy-20106.upstash.io:6379)/go_testing")
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return db, nil
}
