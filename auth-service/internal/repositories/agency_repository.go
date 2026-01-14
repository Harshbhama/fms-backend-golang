package repositories

import (
	"database/sql"
	"github.com/yourusername/auth-service/internal/models"
)

type AgencyRepository struct {
	DB *sql.DB
}

func NewAgencyRepository(db *sql.DB) *AgencyRepository {
	return &AgencyRepository{DB: db}
}

func (r *AgencyRepository) CreateAgency(a *models.Agency) error {
	query := `
	INSERT INTO agencies (
		id,
		name,
		email,
		website,
		description,
		location,
		created_at
	)
	VALUES ($1, $2, $3, $4, $5, $6, NOW())
	RETURNING id
`

	return r.DB.QueryRow(
		query,
		a.ID,
		a.Name,
		a.Email,
		a.Website,
		a.Description,
		a.Location,
	).Scan(&a.ID)
}
