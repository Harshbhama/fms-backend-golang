package repositories

import (
	"database/sql"
	"github.com/lib/pq"
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
		team_size,
		founded_year,
		min_budget,
		avg_hourly_rate,
		specializations,
		services,
		phone,
		address,
		certifications,
		languages,
		send_invitation,
		add_to_favorites
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
`

	_, err := r.DB.Exec(
		query,
		a.ID,
		a.Name,
		a.Email,
		a.Website,
		a.Description,
		a.Location,
		a.TeamSize,
		a.FoundedYear,
		a.MinBudget,
		a.AvgHourlyRate,
		pq.Array(a.Specializations),
		pq.Array(a.Services),
		a.Phone,
		a.Address,
		pq.Array(a.Certifications),
		pq.Array(a.Languages),
		a.SendInvitation,
		a.AddToFavorites,
	)
	return err
}
