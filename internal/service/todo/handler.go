package todo

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Tasks)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTask.ID = uuid.New().String()
	newTask.CreatedAt = time.Now()

	if newTask.Position == 0 {
		newTask.Position = len(Tasks) + 1
	}

	Tasks = append(Tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}
