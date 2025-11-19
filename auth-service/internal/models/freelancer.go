package models



type Freelancer struct {
	ID                int64    `json:"id"`
	 Firstname         *string  `json:"first_name,omitempty"`
    Lastname          *string  `json:"last_name,omitempty"`
    ProfessionalTitle *string  `json:"professional_title,omitempty"`
    ProfessionalBio   *string  `json:"professional_bio,omitempty"`
    Location          *string  `json:"location,omitempty"`
    HourlyRate        *float64 `json:"hourly_rate,omitempty"`
    ExperienceLevel   *string  `json:"experience_level,omitempty"`
    Availability      *string  `json:"availability,omitempty"`
    Skills            []string `json:"skills,omitempty"`
    CreatedAt         *string  `json:"created_at,omitempty"`
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

type FreelancerClientJoin struct {
	FreelancerID int64  `json:"freelancer_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	ClientID     int64  `json:"client_id"`
	CreatedAt		string `json:"created_at"`
	Email				string `json:"email"`
}