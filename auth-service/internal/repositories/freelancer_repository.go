package repositories

import (
	"database/sql"

	"github.com/yourusername/auth-service/internal/models"
)

type FreelancerRepository struct {
	DB *sql.DB
}

func NewFreelancerRepository(db *sql.DB) *FreelancerRepository {
	return &FreelancerRepository{DB: db}
}

func (r *FreelancerRepository) CreateFreelancer(f *models.Freelancer) error {
	query := `INSERT INTO freelancers (id, first_name, last_name, created_at)
		VALUES ($1, $2, $3, NOW()) RETURNING id`

	return r.DB.QueryRow(query, f.ID, f.Firstname, f.Lastname).Scan(&f.ID)
}

func (r *FreelancerRepository) CreateFreelancerRates(fr *models.FreelancerRates) error {
	query := `INSERT INTO freelancer_rates (freelancer_id, rate, currency, created_at)
		VALUES ($1, $2, $3, NOW()) RETURNING id`

	return r.DB.QueryRow(query, fr.FreelancerID, fr.Rate, fr.Currency).Scan(&fr.ID)
}

func (r *FreelancerRepository) CreateFreelancerProject(fp *models.FreelancerProject) error {
	query := `INSERT INTO freelancer_projects (freelancer_id, project_id, created_at)
		VALUES ($1, $2, NOW()) RETURNING freelancer_id, project_id`
	_, err := r.DB.Exec(query, fp.FreelancerId, fp.ProjectId)
	if err != nil {
		return err
	}
	return nil
}