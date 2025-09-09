package http

import (
	"net/http"
	"time"

	"github.com/0x2BadB002/todo-backend/tasks"
)

type Server struct {
	server  *http.Server
	handler *Handler
}

func New(tasks *tasks.Tasks) *Server {
	server := &http.Server{
		Addr:              ":9000",
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	handler := NewHandler(tasks)

	http.HandleFunc("/tasks/create", handler.CreateTask)

	return &Server{
		server:  server,
		handler: handler,
	}
}

func (h *Server) Serve() error {
	return h.server.ListenAndServe()
}
