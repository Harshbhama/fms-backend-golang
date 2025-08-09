package models

import "time"

type ClientFreelancer struct {
	ClientId        int64     `json:"client_id"`
	FreelancerId    int64     `json:"freelancer_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
