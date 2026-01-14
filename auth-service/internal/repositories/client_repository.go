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

func (r *ClientRepository) CreateClientFreelancer(clientFreelancer *models.ClientFreelancer) error {
	query := `INSERT INTO client_freelancers (client_id, freelancer_id, created_at)
	          VALUES ($1, $2, NOW()) RETURNING client_id, freelancer_id`
	_, err := r.db.Exec(query, clientFreelancer.ClientId, clientFreelancer.FreelancerId)
	if err != nil {
		return err
	}
	return nil
}

func (r *ClientRepository) CreateClientAgency(clientAgency *models.ClientAgency) error {
	query := `INSERT INTO client_agencies (client_id, agency_id, created_at)
	          VALUES ($1, $2, NOW()) RETURNING client_id, agency_id`
	_, err := r.db.Exec(query, clientAgency.ClientId, clientAgency.AgencyId)
	if err != nil {
		return err
	}
	return nil
}

func (r *ClientRepository) CreateClientProject(clientProject *models.ClientProject) error {
	query := `INSERT INTO client_projects (client_id, project_id, created_at)
	          VALUES ($1, $2, NOW()) RETURNING client_id, project_id`
	_, err := r.db.Exec(query, clientProject.ClientId, clientProject.ProjectId)
	if err != nil {
		return err
	}
	return nil
}

func (r *ClientRepository) GetProjectsByClientId(id uint) (*[] models.ClientProjects, error) {
	query := `
	SELECT 
		cp.client_id, 
		cp.project_id, 
		COALESCE(p.status, '') AS status,
		COALESCE(p.description, '') AS description,
		COALESCE(p.category, '') AS category,
		COALESCE(p.priority, '') AS priority,
		COALESCE(p.required_skills, '{}') AS required_skills,
		COALESCE(p.custom_skills, '{}') AS custom_skills,
		COALESCE(p.detailed_requirements, '') AS detailed_requirements,
		COALESCE(p.expected_deliverables, '') AS expected_deliverables,
		COALESCE(p.assignment_timing, '') AS assignment_timing,
		COALESCE(p.timeline, '{}'::jsonb) AS timeline
	FROM client_projects cp
	INNER JOIN projects p ON cp.project_id = p.id
	WHERE cp.client_id = $1
`


	rows, err := r.db.Query(query, id)


	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []models.ClientProjects
	
	for rows.Next() {
		var project models.ClientProjects
		err := rows.Scan(&project.ClientID, &project.ProjectID, &project.Status, &project.Description, &project.Category,
			&project.Priority, &project.RequiredSkills, &project.CustomSkills, &project.DetailedRequirements,
			&project.ExpectedDeliverables, &project.AssignmentTiming, &project.Timeline)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &projects, nil
	
}