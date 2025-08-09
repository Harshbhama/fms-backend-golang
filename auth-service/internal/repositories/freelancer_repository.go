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
