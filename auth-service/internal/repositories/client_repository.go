package repositories

import (
	"database/sql"
	// "time"
	// "fmt"
	// "github.com/aws/aws-sdk-go-v2/aws/protocol/query"
	"github.com/yourusername/auth-service/internal/models"
	// "github.com/yourusername/auth-service/internal/utils"
	
	// "errors"
	
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) CreateClient(client *models.Client) error {

	
	query := `INSERT INTO clients (id, first_name, last_name, created_at, updated_at)
	          VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`

	return r.db.QueryRow(query, client.ID, client.FirstName, client.LastName).Scan(&client.ID)
}

