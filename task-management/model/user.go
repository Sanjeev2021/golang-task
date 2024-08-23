package model

import "time"

type User struct {
	ID        uint      `json: "id"`
	Name      string    `json: "name"`
	Email     string    `json: "email"`
	Password  string    `json: "email"`
	CreatedAt time.Time `json: "createdat"`
	UpdatedAt time.Time `json: "updatedat"`
}
