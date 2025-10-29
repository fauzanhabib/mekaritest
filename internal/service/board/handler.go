package board

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func GetBoards(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	w.Header().Set("Content-Type", "application/json")

	if userID == "" {
		json.NewEncoder(w).Encode(Boards)
		return
	}

	filtered := []Board{}
	for _, b := range Boards {
		if b.OwnerUserID == userID {
			filtered = append(filtered, b)
		}
	}
	json.NewEncoder(w).Encode(filtered)
}

func AddBoard(w http.ResponseWriter, r *http.Request) {
	var b Board
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b.ID = uuid.New().String()
	b.CreatedAt = time.Now()
	Boards = append(Boards, b)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(b)
}
