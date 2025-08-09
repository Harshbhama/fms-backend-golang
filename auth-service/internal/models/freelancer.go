package models

type Freelancer struct {
	ID    int64 `json:"id"`
	Firstname  string `json:"first_name"`
	Lastname string `json:"last_name"`
	CreatedAt string `json:"created_at"`
}

