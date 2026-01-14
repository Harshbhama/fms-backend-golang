package models

import "github.com/lib/pq"

type Agency struct {
	ID              int64          `json:"id"`
	Name            string         `json:"name" binding:"required,min=1,max=100"`
	Email           string         `json:"email" binding:"required,email"`
	Website         *string        `json:"website,omitempty"`
	Description     string         `json:"description" binding:"required,min=50,max=1000"`
	Location        string         `json:"location" binding:"required"`
	TeamSize        int            `json:"team_size" binding:"required,min=1"`
	FoundedYear     int            `json:"founded_year" binding:"required,min=1900"`
	MinBudget       float64        `json:"min_budget" binding:"required,min=1000"`
	AvgHourlyRate   float64        `json:"avg_hourly_rate" binding:"required,min=10"`
	Specializations pq.StringArray `json:"specializations" binding:"required,min=1"`
	Services        pq.StringArray `json:"services" binding:"required,min=1"`
	Phone           *string        `json:"phone,omitempty"`
	Address         *string        `json:"address,omitempty"`
	Certifications  pq.StringArray `json:"certifications,omitempty"`
	Languages       pq.StringArray `json:"languages,omitempty"`
	SendInvitation  bool           `json:"send_invitation"`
	AddToFavorites  bool           `json:"add_to_favorites"`
	CreatedAt       *string        `json:"created_at,omitempty"`
}

type AgencyClientJoin struct {
	AgencyID        int64          `json:"agency_id"`
	Name            string         `json:"name"`
	Email           string         `json:"email"`
	Website         *string        `json:"website,omitempty"`
	Description     string         `json:"description"`
	Location        string         `json:"location"`
	TeamSize        int            `json:"team_size"`
	FoundedYear     int            `json:"founded_year"`
	MinBudget       float64        `json:"min_budget"`
	AvgHourlyRate   float64        `json:"avg_hourly_rate"`
	Specializations pq.StringArray `json:"specializations"`
	Services        pq.StringArray `json:"services"`
	Phone           *string        `json:"phone,omitempty"`
	Address         *string        `json:"address,omitempty"`
	Certifications  pq.StringArray `json:"certifications,omitempty"`
	Languages       pq.StringArray `json:"languages,omitempty"`
	ClientID        int64          `json:"client_id"`
	CreatedAt       string         `json:"created_at"`
}
