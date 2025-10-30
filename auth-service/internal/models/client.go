package models

import( 
	"time"
	"github.com/lib/pq"
	"encoding/json"
)

type Client struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}

type ClientProjects struct {
	ClientID             uint           `json:"client_id"`
	ProjectID            uint           `json:"project_id"`
	Status               string         `json:"status"`
	Description          string         `json:"description"`
	Category             string         `json:"category"`
	Priority             string         `json:"priority"`
	RequiredSkills       pq.StringArray `json:"required_skills" db:"required_skills"`
	CustomSkills         pq.StringArray `json:"custom_skills" db:"custom_skills"`
	DetailedRequirements string         `json:"detailed_requirements"`
	ExpectedDeliverables string         `json:"expected_deliverables"`
	AssignmentTiming     string         `json:"assignment_timing"`
	Timeline             json.RawMessage `json:"timeline"` // holds any valid JSON object
}
