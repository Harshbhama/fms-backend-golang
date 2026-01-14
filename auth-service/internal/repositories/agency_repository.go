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
		location
	)
	VALUES ($1, $2, $3, $4, $5, $6)
`

	_, err := r.DB.Exec(
		query,
		a.ID,
		a.Name,
		a.Email,
		a.Website,
		a.Description,
		a.Location,
	)
	return err
}
