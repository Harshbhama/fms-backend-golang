package repositories

import (
	"database/sql"
	"time"
	"fmt"
	// "github.com/aws/aws-sdk-go-v2/aws/protocol/query"
	"github.com/yourusername/auth-service/internal/models"
	"github.com/yourusername/auth-service/internal/utils"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`

	return r.db.QueryRow(query, user.Email, user.Password, user.Role, time.Now(), time.Now()).Scan(&user.ID)
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
	query := `UPDATE users SET Email = $1, Password = $2, Role = $3 WHERE id = $4`

	_, err := r.db.Exec(query, user.Email, user.Password, user.Role, id)
	return err
}

func (r *UserRepository) LoginUser(loginInput *models.UserLogin) (*models.User, error) {
	query := `Select id, email, password, role, created_at, updated_at from users where email = $1`

	row := r.db.QueryRow(query, loginInput.Email)
	
	var users models.User
	err := row.Scan(&users.ID, &users.Email, &users.Password, &users.Role, &users.CreatedAt, &users.UpdatedAt)
	if err == sql.ErrNoRows {
		print("email not found")
		return nil, errors.New("email not found")
	}
	fmt.Println("DB hash:", users.Password)
	fmt.Println("Candidate pwd:", loginInput.Password)
	if(utils.CheckPasswordHash(loginInput.Password, users.Password)){
		if err != nil {
			return nil, err
		}
		print("here")
		return &users, nil
	} else{
		return nil, errors.New("password is not correct")
	}

}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, email, password, role, created_at, updated_at FROM users WHERE email = $1`

	row := r.db.QueryRow(query, email)

	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}