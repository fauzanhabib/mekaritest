package router

import (
	"go-board-app/internal/service/board"
	"go-board-app/internal/service/todo"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Board routes
	r.HandleFunc("/boards", board.GetBoards).Methods("GET", "OPTIONS")
	r.HandleFunc("/boards", board.AddBoard).Methods("POST", "OPTIONS")

	// Todo routes
	r.HandleFunc("/todos", todo.GetTodos).Methods("GET")
	r.HandleFunc("/todos", todo.AddTodo).Methods("POST")

	return r
}
