package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/0x2BadB002/todo-backend/domain"
	"github.com/0x2BadB002/todo-backend/tasks"
)

type Handler struct {
	tasks *tasks.Tasks
}

func NewHandler(tasks *tasks.Tasks) *Handler {
	return &Handler{tasks: tasks}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req domain.AddTaskRequest
	err := decoder.Decode(&req)
	if err != nil {
		log.Println(err)
		return
	}

	id, err := h.tasks.CreateTask(context.Background(), req)
	if err != nil {
		log.Println(err)
		return
	}

	resp := map[string]interface{}{"id": id}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("encode response:", err)
	}
}
