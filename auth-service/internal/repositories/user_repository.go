package repositories

import (
	"database/sql"
	"time"

	"github.com/yourusername/auth-service/internal/models"

)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4) RETURNING id`

	return r.db.QueryRow(query, user.Email, user.Password, time.Now(), time.Now()).Scan(&user.ID)
}

func (r *UserRepository) GetUser(id int64) (*models.User, error) {
	query := `SELECT * FROM users WHERE id = $1`

	row := r.db.QueryRow(query, id)

	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(id int64, user *models.User) error {
	query := `UPDATE users SET first_name = $1, last_name = $2, email = $3, password = $4, updated_at = $5 WHERE id = $6`

	_, err := r.db.Exec(query, user.Email, user.Password, time.Now(), id)
	return err
}
	