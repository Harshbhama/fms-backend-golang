package repositories

import (
	"database/sql"

	"github.com/yourusername/auth-service/internal/models"
	"encoding/json"
	"github.com/lib/pq"
)

type ProjectRepository struct {
	DB *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (r *ProjectRepository) CreateProject(p *models.Projects) error {
	query := `
	INSERT INTO projects (
		name,
		description,
		status,
		category,
		priority,
		required_skills,
		custom_skills,
		detailed_requirements,
		expected_deliverables,
		assignment_timing,
		timeline,
		created_at
	)
	VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW()
	)
	RETURNING id`

	// Convert struct fields to suitable types for PostgreSQL
	var timelineJSON any
	if p.Timeline != nil {
		timelineJSON, _ = json.Marshal(p.Timeline)
	}
 
	return r.DB.QueryRow(
		query,
		p.Name,
		p.Description,
		p.Status,
		p.Category,
		p.Priority,
		pq.Array(p.RequiredSkills), // ✅ fixed
		pq.Array(p.CustomSkills),   // ✅ fixed
		p.DetailedRequirements,
		p.ExpectedDeliverables,
		p.AssignmentTiming,
		timelineJSON,
	).Scan(&p.ID)
}
