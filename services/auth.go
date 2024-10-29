package services

import (
	"database/sql"
	"errors"
	"rngdev/models"
)

func Register(db *sql.DB, user models.User) error {
	query := "INSERT INTO users (number,email,password) VALUES (?,?,?)"
	_, err := db.Exec(query, user.Number, user.Email, user.Password)
	return err
}

func Login(db *sql.DB, number, password string) error {
	var user models.User
	query := "SELECT id, number, email from users where number = ? and password = ?"
	err := db.QueryRow(query, number, password).Scan(&user.ID, &user.Number, &user.Email)
	if err != nil {
		return errors.New("user not found")
	}
	return nil
}
