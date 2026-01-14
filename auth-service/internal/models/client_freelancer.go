package models

import "time"

type ClientFreelancer struct {
	ClientId     int64     `json:"client_id"`
	FreelancerId int64     `json:"freelancer_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ClientProject struct {
	ClientId  int64     `json:"client_id"`
	ProjectId int64     `json:"project_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ClientAgency struct {
	ClientId  int64     `json:"client_id"`
	AgencyId  int64     `json:"agency_id"`
	CreatedAt time.Time `json:"created_at"`
}
