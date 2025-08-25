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

type FreelancerTimesheet struct {
	ID        int64   `json:"id"`
	FreelancerID int64 `json:"freelancer_id"`
	ProjectID int64 `json:"project_id"`
}

type FreelancerTimesheetMetadata struct {
	MetadataID int64  `json:"metadata_id"`
	TimesheetID int64  `json:"timesheet_id"`
	Date 		 string `json:"date"`
	Hours				float64 `json:"hours"`
	Status			string  `json:"status"`
	Remarks			string  `json:"remarks"`
	CreatedAt   string  `json:"created_at"`
}
