package db

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/0x2BadB002/todo-backend/domain"
)

var (
	//go:embed create_table.sql
	create_table_query string
	//go:embed insert_task.sql
	insert_query string
	//go:embed get_tasks.sql
	get_query string
)

var (
	ErrInsertIntoDB = errors.New("failed inserting task to db")
	ErrGetFromDB    = errors.New("failed retrieving tasks from db")
)

type DB struct {
	filename string

	db *sql.DB
}

func New(filename string) *DB {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(create_table_query)
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		filename: filename,
		db:       db,
	}
}

func (db *DB) Close() {
	err := db.db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB) AddTask(ctx context.Context, req domain.AddTaskRequest) (int64, error) {
	res, err := db.db.ExecContext(
		ctx,
		insert_query,
		req.Name,
		req.Description,
		req.Priority,
		req.Due,
	)
	if err != nil {
		return -1, errors.Join(ErrInsertIntoDB, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, errors.Join(ErrInsertIntoDB, err)
	}

	return id, nil
}

func (db *DB) GetTasks(ctx context.Context) (res []domain.Task, err error) {
	rows, err := db.db.QueryContext(ctx, get_query)
	if err != nil {
		return nil, errors.Join(ErrGetFromDB, err)
	}
	defer func() {
		tmpErr := rows.Close()
		if tmpErr != nil && err != nil {
			err = errors.Join(err, tmpErr)
			return
		}
		if tmpErr != nil {
			err = errors.Join(ErrGetFromDB, tmpErr)
			return
		}
	}()

	res = []domain.Task{}
	for rows.Next() {
		var task domain.Task

		if rows.Err() != nil {
			return res, errors.Join(ErrGetFromDB, err)
		}

		err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Description,
			&task.Priority,
			&task.Due,
			&task.UpdatedAt,
			&task.CreatedAt,
			&task.DeletedAt,
		)
		if err != nil {
			return res, errors.Join(ErrGetFromDB, err)
		}

		res = append(res, task)
	}
	if rows.Err() != nil {
		return res, errors.Join(ErrGetFromDB, err)
	}

	return res, nil
}
