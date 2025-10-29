package todo

import "time"

type Task struct {
	ID          string     `json:"id"`
	ListID      string     `json:"list_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date"`
	Position    int        `json:"position"`
	CreatedAt   time.Time  `json:"created_at"`
}
