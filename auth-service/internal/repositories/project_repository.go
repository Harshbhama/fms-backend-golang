package repositories

import (
	"database/sql"

	"github.com/yourusername/auth-service/internal/models"
)

type ProjectRepository struct {
	DB *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (r *ProjectRepository) CreateProject(p *models.Projects) error {
	query := `INSERT INTO projects (name, description, status, created_at)
		VALUES ($1, $2, $3, NOW()) RETURNING id`

	return r.DB.QueryRow(query, p.Name, p.Description, p.Status).Scan(&p.ID)
}