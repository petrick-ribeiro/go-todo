package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/petrick-ribeiro/go-todo/storage"
)

type APIServer struct {
	listenAddr string
	storage    storage.Storage
}

func NewAPIServer(listenAddr string, db storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		storage:    db,
	}
}

func (s *APIServer) Run() {
	r := mux.NewRouter()

	// Handlers
	r.HandleFunc("/todo", makeHTTPHandleFunc(s.handleTodo))
	r.HandleFunc("/todo/{id}", makeHTTPHandleFunc(s.handleTodoByID))

	log.Printf("Starting the server on port %s", s.listenAddr)

	http.ListenAndServe(s.listenAddr, r)
}
