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
)

var (
	ErrDBInsert = errors.New("failed inserting task to db")
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
		req.Due_at,
	)
	if err != nil {
		return -1, errors.Join(err, ErrDBInsert)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, errors.Join(err, ErrDBInsert)
	}

	return id, nil
}
