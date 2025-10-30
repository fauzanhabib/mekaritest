package board

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

// --- ADD THIS ENTIRE FUNCTION ---
func UpdateBoard(w http.ResponseWriter, r *http.Request) {
	// 1. Get the ID from the URL path (e.g., /api/boards/xxxx-xxxx)
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Missing board ID", http.StatusBadRequest)
		return
	}

	// 2. Decode the incoming JSON body into an "updatedBoard" struct
	var updatedBoard Board
	if err := json.NewDecoder(r.Body).Decode(&updatedBoard); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 3. Find the board in our in-memory slice
	found := false
	var boardToReturn Board
	for i, b := range Boards {
		if b.ID == id {
			// 4. Update all fields with the new data
			Boards[i].Name = updatedBoard.Name
			Boards[i].Description = updatedBoard.Description
			Boards[i].OwnerUserID = updatedBoard.OwnerUserID
			Boards[i].Status = updatedBoard.Status
			Boards[i].Priority = updatedBoard.Priority
			Boards[i].Labels = updatedBoard.Labels
			Boards[i].DueDate = updatedBoard.DueDate
			// We keep the original ID and CreatedAt time

			boardToReturn = Boards[i] // Get the final updated board to send back
			found = true
			break
		}
	}

	// 5. Handle if not found
	if !found {
		http.Error(w, "Board not found", http.StatusNotFound)
		return
	}

	// 6. Respond with the updated board
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(boardToReturn)
}

func DeleteBoard(w http.ResponseWriter, r *http.Request) {
	// 1. Get the ID from the URL path
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Missing board ID", http.StatusBadRequest)
		return
	}

	// 2. Find the board and remove it
	found := false
	for i, b := range Boards {
		if b.ID == id {
			// Remove the item from the slice
			// This syntax appends the slice before the item...
			// with the slice after the item.
			Boards = append(Boards[:i], Boards[i+1:]...)
			found = true
			break
		}
	}

	// 3. Handle if not found
	if !found {
		http.Error(w, "Board not found", http.StatusNotFound)
		return
	}

	// 4. Respond with success
	w.WriteHeader(http.StatusNoContent) // 204 No Content is a standard for successful deletes
}
