package model

import "time"

type Task struct {
	ID          uint      `json:"id"`
	UserID      uint      `json: "userid"`
	User        User      `json: "user" gorm: foreignKey:UserID"`
	Title       string    `json:"title"`
	Description string    `json:"Description"`
	Priority    string    `json:"Priority"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json: "createdat"`
	UpdatedAt   time.Time `json: "updatedat"`
	DueDate     time.Time `json: "duedate`
}
