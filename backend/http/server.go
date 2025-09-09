package http

import (
	"net/http"
	"time"

	v1 "github.com/0x2BadB002/todo-backend/http/v1"
	"github.com/0x2BadB002/todo-backend/tasks"
)

type Server struct {
	server *http.Server

	handler *v1.Handler
}

func New(tasks *tasks.Tasks) *Server {
	handler := v1.New(tasks)

	mux := create_mux(handler)

	server := &http.Server{
		Addr:              ":9000",
		Handler:           mux,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	return &Server{
		server:  server,
		handler: handler,
	}
}

func (h *Server) Serve() error {
	return h.server.ListenAndServe()
}

func create_mux(handler *v1.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/tasks/new", handler.CreateTask)

	return mux
}
