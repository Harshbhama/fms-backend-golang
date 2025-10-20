package models

import "time"

type FreelancerProject struct {
	FreelancerIds []int64     `json:"freelancer_id"`
	ProjectId    int64     `json:"project_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}


