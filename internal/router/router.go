package router

import (
	"go-board-app/internal/service/board"
	"go-board-app/internal/service/todo"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	// Board routes
	api.HandleFunc("/boards", board.GetBoards).Methods("GET", "OPTIONS")
	api.HandleFunc("/boards", board.AddBoard).Methods("POST", "OPTIONS")
	api.HandleFunc("/boards/{id}", board.UpdateBoard).Methods("PUT", "OPTIONS")
	api.HandleFunc("/boards/{id}", board.DeleteBoard).Methods("DELETE", "OPTIONS")

	// Todo routes (ini kemarin yg dibuat test)
	api.HandleFunc("/todos", todo.GetTodos).Methods("GET")
	api.HandleFunc("/todos", todo.AddTodo).Methods("POST")

	return r
}
