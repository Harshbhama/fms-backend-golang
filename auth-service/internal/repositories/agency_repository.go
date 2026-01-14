package repositories

import (
	"database/sql"
	"fmt"

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
	INSERT INTO agency (
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

func (r *AgencyRepository) GetAgencyByClientID(id uint, search string) (*[]models.AgencyClientJoin, error) {
	// Note: For optimal performance, ensure database indexes exist on:
	// - client_agencies.client_id
	// - agency(name, email, location, description)
	query := `SELECT ca.agency_id, ag.name, ag.email, ag.website, ag.description, ag.location, 
		ag.team_size, ag.founded_year, ag.min_budget, ag.avg_hourly_rate, ag.specializations, 
		ag.services, ag.phone, ag.address, ag.certifications, ag.languages, ca.client_id, ca.created_at
		FROM client_agencies ca
		INNER JOIN agency ag 
			ON ag.id = ca.agency_id
		WHERE ca.client_id = $1`

	var rows *sql.Rows
	var err error

	if search != "" {
		// Validate search input length to prevent abuse
		if len(search) > 100 {
			return nil, fmt.Errorf("search query too long (max 100 characters)")
		}
		query += ` AND (
			LOWER(ag.name) LIKE LOWER($2) OR 
			LOWER(ag.email) LIKE LOWER($2) OR 
			LOWER(ag.location) LIKE LOWER($2) OR 
			LOWER(ag.description) LIKE LOWER($2) OR
			EXISTS (
				SELECT 1 FROM unnest(ag.specializations) AS spec 
				WHERE LOWER(spec) LIKE LOWER($2)
			) OR
			EXISTS (
				SELECT 1 FROM unnest(ag.services) AS service 
				WHERE LOWER(service) LIKE LOWER($2)
			)
		)`
		// Safe: searchPattern is passed as parameterized query argument, preventing SQL injection
		searchPattern := "%" + search + "%"
		rows, err = r.DB.Query(query, id, searchPattern)
	} else {
		rows, err = r.DB.Query(query, id)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agencies []models.AgencyClientJoin

	for rows.Next() {
		var a models.AgencyClientJoin
		if err := rows.Scan(&a.AgencyID, &a.Name, &a.Email, &a.Website, &a.Description,
			&a.Location, &a.TeamSize, &a.FoundedYear, &a.MinBudget, &a.AvgHourlyRate,
			&a.Specializations, &a.Services, &a.Phone, &a.Address,
			&a.Certifications, &a.Languages, &a.ClientID, &a.CreatedAt); err != nil {
			return nil, err
		}
		agencies = append(agencies, a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(agencies) == 0 {
		return nil, sql.ErrNoRows
	}

	return &agencies, nil
}
