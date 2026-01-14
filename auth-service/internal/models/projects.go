package models

type Projects struct {
	ID                   int64     `json:"id"`
	Name                 string    `json:"name"`
	Description          string    `json:"description"`
	Status               string    `json:"status"`
	CreatedAt            string    `json:"created_at"`
	Category             *string   `json:"category,omitempty"`
	Priority             *string   `json:"priority,omitempty"`
	RequiredSkills       *[]string `json:"required_skills,omitempty"`
	CustomSkills         *[]string `json:"custom_skills,omitempty"`
	DetailedRequirements *string   `json:"detailed_requirements,omitempty"`
	ExpectedDeliverables *string   `json:"expected_deliverables,omitempty"`
	AssignmentTiming     *string   `json:"assignment_timing,omitempty"`
	Timeline             *struct {
		Type              *string `json:"type,omitempty"`
		StartDate         *string `json:"startDate,omitempty"`
		EndDate           *string `json:"endDate,omitempty"`
		EstimatedDuration *int    `json:"estimatedDuration,omitempty"`
	} `json:"timeline,omitempty"`
}
