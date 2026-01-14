package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/yourusername/auth-service/internal/models"
	"github.com/yourusername/auth-service/internal/utils"
)

type FreelancerRepository struct {
	DB *sql.DB
}

func NewFreelancerRepository(db *sql.DB) *FreelancerRepository {
	return &FreelancerRepository{DB: db}
}

func (r *FreelancerRepository) CreateFreelancer(f *models.Freelancer) error {
	query := `
	INSERT INTO freelancers (
		id,
		first_name,
		last_name,
		professional_title,
		professional_bio,
		location,
		hourly_rate,
		experience_level,
		availability,
		skills,
		created_at
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW())
	RETURNING id
`

	return r.DB.QueryRow(
		query,
		f.ID,
		f.Firstname,
		f.Lastname,
		f.ProfessionalTitle,
		f.ProfessionalBio,
		f.Location,
		f.HourlyRate,
		f.ExperienceLevel,
		f.Availability,
		pq.Array(f.Skills), // use pq.Array for []string
	).Scan(&f.ID)
}

func (r *FreelancerRepository) CreateFreelancerRates(fr *models.FreelancerRates) error {
	query := `INSERT INTO freelancer_rates (freelancer_id, rate, currency, created_at)
		VALUES ($1, $2, $3, NOW()) RETURNING id`

	return r.DB.QueryRow(query, fr.FreelancerID, fr.Rate, fr.Currency).Scan(&fr.ID)
}

func (r *FreelancerRepository) CreateFreelancerProject(fp *models.FreelancerProject) error {
	query := `INSERT INTO freelancer_projects (freelancer_id, project_id, created_at)
	          VALUES ($1, $2, NOW())`

	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, freelancerID := range fp.FreelancerIds {
		_, err := tx.Exec(query, freelancerID, fp.ProjectId)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *FreelancerRepository) CreateFreelancerTimesheet(ft *models.FreelancerTimesheet) error {
	query := `INSERT INTO freelancer_timesheet (freelancer_id, project_id, created_at)
		VALUES ($1, $2, NOW()) RETURNING id`

	return r.DB.QueryRow(query, ft.FreelancerID, ft.ProjectID).Scan(&ft.ID)
}

func (r *FreelancerRepository) CreateFreelancerTimesheetMetadata(ftm *models.FreelancerTimesheetMetadata) error {

	table := utils.GetSharedMetadataTableName(ftm.TimesheetID)

	query := fmt.Sprintf(`
		INSERT INTO %s (timesheet_id, date, hours, status, remarks, created_at)
		VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING metadata_id
	`, table)

	return r.DB.QueryRow(query, ftm.TimesheetID, ftm.Date, ftm.Hours, ftm.Status, ftm.Remarks).Scan(&ftm.MetadataID)
}

func (r *FreelancerRepository) GetFreelancerByClientID(id uint) (*[]models.FreelancerClientJoin, error) {
	query := `SELECT freelancer_id, first_name, last_name, fl.created_at, email, client_id
		FROM client_freelancers cl
		INNER JOIN freelancers fl 
			ON fl.id = cl.freelancer_id
		INNER JOIN users u 
			ON u.id = fl.id
		WHERE client_id = $1`

	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var freelancers []models.FreelancerClientJoin

	for rows.Next() {
		var f models.FreelancerClientJoin
		if err := rows.Scan(&f.FreelancerID, &f.FirstName, &f.LastName, &f.CreatedAt, &f.Email, &f.ClientID); err != nil {
			return nil, err
		}
		freelancers = append(freelancers, f)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(freelancers) == 0 {
		return nil, errors.New("no rows found")
	}

	return &freelancers, nil
}
