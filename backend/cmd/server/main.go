package main

import (
	"log"

	"github.com/0x2BadB002/todo-backend/db"
	"github.com/0x2BadB002/todo-backend/http"
	"github.com/0x2BadB002/todo-backend/tasks"
)

func main() {
	db_name := "tasks.db"

	db := db.New(db_name)
	defer db.Close()

	tasks := tasks.New(db)

	http := http.New(tasks)

	err := http.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
