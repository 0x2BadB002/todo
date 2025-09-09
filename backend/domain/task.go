package domain

import (
	"time"
)

type Task struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Priority    int        `json:"priority"`
	Due         *time.Time `json:"due"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type AddTaskRequest struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Priority    int        `json:"priority"`
	Due         *time.Time `json:"due"`
}
