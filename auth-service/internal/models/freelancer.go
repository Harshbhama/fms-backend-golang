package models

type Freelancer struct {
	ID    int64 `json:"id"`
	Firstname  string `json:"first_name"`
	Lastname string `json:"last_name"`
	CreatedAt string `json:"created_at"`
}

type FreelancerRates struct {
	ID        int64   `json:"id"`
	FreelancerID int64 `json:"freelancer_id"`
	Rate      float64 `json:"rate"`
	Currency  string  `json:"currency"`
	CreatedAt string  `json:"created_at"`
}