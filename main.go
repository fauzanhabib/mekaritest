package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// --- Structs based on ERD ---

type User struct {
	ID           string    `json:"id" gorm:"type:uuid;primary_key"`
	Username     string    `json:"username" gorm:"username"`
	Email        string    `json:"email" gorm:"unique;not null"`
	PasswordHash string    `json:"-" gorm:"password_hash;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
}

type Board struct {
	ID          string    `json:"id" gorm:"type:uuid;primary_key"`
	Name        string    `json:"name" gorm:"name;not null"`
	Description string    `json:"description" gorm:"description"`
	OwnerUserID string    `json:"owner_user_id" gorm:"type:uuid;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
}

type List struct {
	ID       string `json:"id" gorm:"type:uuid;primary_key"`
	BoardID  string `json:"board_id" gorm:"type:uuid;not null"`
	Name     string `json:"name" gorm:"name;not null"`
	Position int    `json:"position" gorm:"position;not null"`
}

type Task struct {
	ID          string     `json:"id" gorm:"type:uuid;primary_key"`
	ListID      string     `json:"list_id" gorm:"type:uuid;not null"`
	Title       string     `json:"title" gorm:"title;not null"`
	Description string     `json:"description" gorm:"description"`
	DueDate     *time.Time `json:"due_date" gorm:"due_date"`
	Position    int        `json:"position" gorm:"position;not null"`
	CreatedAt   time.Time  `json:"created_at" gorm:"created_at"`
}

type BoardMember struct {
	UserID  string `json:"user_id" gorm:"type:uuid;primary_key"`
	BoardID string `json:"board_id" gorm:"type:uuid;primary_key"`
	Role    string `json:"role" gorm:"role;default:'member'"`
}

type TaskAssignee struct {
	UserID string `json:"user_id" gorm:"type:uuid;primary_key"`
	TaskID string `json:"task_id" gorm:"type:uuid;primary_key"`
}

// --- In-memory storage (temporary) ---

var tasks []Task
var boards []Board

// --- Handlers ---

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	var newTask Task

	// Decode JSON body
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate UUID and timestamps
	newTask.ID = uuid.New().String()
	newTask.CreatedAt = time.Now()

	// Default position if not provided
	if newTask.Position == 0 {
		newTask.Position = len(tasks) + 1
	}

	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

// --- CORS Middleware ---
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// --- Handlers ---
func getBoards(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	w.Header().Set("Content-Type", "application/json")

	if userID == "" {
		json.NewEncoder(w).Encode(boards)
		return
	}

	filtered := []Board{}
	for _, b := range boards {
		if b.OwnerUserID == userID {
			filtered = append(filtered, b)
		}
	}
	json.NewEncoder(w).Encode(filtered)
}

func addBoard(w http.ResponseWriter, r *http.Request) {
	var b Board
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b.ID = uuid.New().String()
	b.CreatedAt = time.Now()
	boards = append(boards, b)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(b)
}

// --- Main ---

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/todos", getTodos).Methods("GET")
	r.HandleFunc("/todos", addTodo).Methods("POST")
	// Routes
	r.HandleFunc("/boards", getBoards).Methods("GET", "OPTIONS")
	r.HandleFunc("/boards", addBoard).Methods("POST", "OPTIONS")

	// Wrap router with CORS middleware
	handler := enableCORS(r)

	fmt.Println("✅ Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
