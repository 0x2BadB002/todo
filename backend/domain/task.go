package domain

import (
	"time"
)

type Task struct {
	ID          int64
	Name        string
	Description string
	Priority    int
	Due_at      time.Time
	Created_at  time.Time
	Updated_at  time.Time
	Deleted_at  time.Time
}

type AddTaskRequest struct {
	Name        string
	Description string
	Priority    int
	Due_at      time.Time
}
