package models

type Agency struct {
	ID          int64   `json:"id"`
	Name        *string `json:"name,omitempty"`
	Email       *string `json:"email,omitempty"`
	Website     *string `json:"website,omitempty"`
	Description *string `json:"description,omitempty"`
	Location    *string `json:"location,omitempty"`
	CreatedAt   *string `json:"created_at,omitempty"`
}
