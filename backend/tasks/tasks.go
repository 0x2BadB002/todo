package tasks

import (
	"context"
	"errors"

	"github.com/0x2BadB002/todo-backend/db"
	"github.com/0x2BadB002/todo-backend/domain"
)

var (
	ErrInvalidRequest = errors.New("failed to handle request")
)

type Tasks struct {
	db *db.DB
}

func New(db *db.DB) *Tasks {
	return &Tasks{
		db: db,
	}
}

func (t *Tasks) CreateTask(ctx context.Context, req domain.AddTaskRequest) (int64, error) {
	id, err := t.db.AddTask(ctx, req)
	if err != nil {
		return -1, errors.Join(ErrInvalidRequest, err)
	}

	return id, nil
}

func (t *Tasks) GetTasks(ctx context.Context) ([]domain.Task, error) {
	tasks, err := t.db.GetTasks(ctx)
	if err != nil {
		return tasks, errors.Join(ErrInvalidRequest, err)
	}

	return tasks, nil
}
