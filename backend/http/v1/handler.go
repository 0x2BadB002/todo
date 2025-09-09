package v1

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

func New(tasks *tasks.Tasks) *Handler {
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
	}

	resp := map[string]any{
		"id": id,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("encode response:", err)
	}
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.tasks.GetTasks(context.Background())
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Println("encode response:", err)
	}
}
