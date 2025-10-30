package board

import "time"

type Board struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerUserID string    `json:"owner_user_id"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	Labels      []string  `json:"labels"`
	DueDate     string    `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
}
