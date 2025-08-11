package models

type Projects struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status     int64  `json:"status"`
	CreatedAt   string `json:"created_at"`
}