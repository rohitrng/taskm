package models

type User struct {
	ID       int    `json:"id"`
	Number   int    `json:"number"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
