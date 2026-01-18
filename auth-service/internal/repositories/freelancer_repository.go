package repositories

import (
	"database/sql"
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

func (r *FreelancerRepository) GetFreelancerByClientID(id uint, search string) (*[]models.FreelancerClientJoin, error) {
	// Note: For optimal performance, ensure database indexes exist on:
	// - client_freelancers.client_id
	// - freelancers(first_name, last_name, location, professional_title)
	// - users.email
	query := `SELECT cl.freelancer_id, fl.first_name, fl.last_name, fl.professional_title, 
		fl.professional_bio, fl.location, fl.hourly_rate, fl.experience_level, 
		fl.availability, fl.skills, cl.client_id, fl.created_at, u.email
		FROM client_freelancers cl
		INNER JOIN freelancers fl 
			ON fl.id = cl.freelancer_id
		INNER JOIN users u 
			ON u.id = fl.id
		WHERE client_id = $1`

	var rows *sql.Rows
	var err error

	if search != "" {
		// Validate search input length to prevent abuse
		if len(search) > 100 {
			return nil, fmt.Errorf("search query too long (max 100 characters)")
		}
		query += ` AND (
			LOWER(fl.first_name) LIKE LOWER($2) OR 
			LOWER(fl.last_name) LIKE LOWER($2) OR 
			LOWER(u.email) LIKE LOWER($2) OR 
			LOWER(fl.professional_title) LIKE LOWER($2) OR 
			LOWER(fl.location) LIKE LOWER($2) OR
			EXISTS (
				SELECT 1 FROM unnest(fl.skills) AS skill 
				WHERE LOWER(skill) LIKE LOWER($2)
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

	var freelancers []models.FreelancerClientJoin

	for rows.Next() {
		var f models.FreelancerClientJoin
		if err := rows.Scan(&f.FreelancerID, &f.FirstName, &f.LastName, &f.ProfessionalTitle, 
			&f.ProfessionalBio, &f.Location, &f.HourlyRate, &f.ExperienceLevel, 
			&f.Availability, pq.Array(&f.Skills), &f.ClientID, &f.CreatedAt, &f.Email); err != nil {
			return nil, err
		}
		freelancers = append(freelancers, f)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(freelancers) == 0 {
		return nil, sql.ErrNoRows
	}

	return &freelancers, nil
}

func (r *FreelancerRepository) GetProjectsByFreelancerID(id uint, search string) (*[]models.FreelancerProjectJoin, error) {
	query := `SELECT fl.freelancer_id, pj.id as project_id, pj.name, pj.description, pj.status,
		pj.category, pj.priority, pj.detailed_requirements, pj.expected_deliverables, 
		pj.assignment_timing, pj.created_at
		FROM freelancer_projects fl
		INNER JOIN projects pj 
			ON fl.project_id = pj.id
		WHERE freelancer_id = $1`

	var rows *sql.Rows
	var err error

	if search != "" {
		// Validate search input length to prevent abuse
		if len(search) > 100 {
			return nil, fmt.Errorf("search query too long (max 100 characters)")
		}
		query += ` AND (
			LOWER(pj.name) LIKE LOWER($2) OR 
			LOWER(pj.description) LIKE LOWER($2) OR 
			LOWER(pj.status) LIKE LOWER($2) OR 
			LOWER(pj.category) LIKE LOWER($2) OR 
			LOWER(pj.priority) LIKE LOWER($2)
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

	var projects []models.FreelancerProjectJoin

	for rows.Next() {
		var p models.FreelancerProjectJoin
		if err := rows.Scan(&p.FreelancerID, &p.ProjectID, &p.Name, &p.Description, &p.Status,
			&p.Category, &p.Priority, &p.DetailedRequirements, &p.ExpectedDeliverables,
			&p.AssignmentTiming, &p.CreatedAt); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &projects, nil
}

func (r *FreelancerRepository) GetClientsByFreelancerID(id uint, search string) (*[]models.ClientFreelancerJoin, error) {
	query := `SELECT cl.client_id, c.first_name, c.last_name, u.email, 
		c.created_at, cl.freelancer_id
		FROM client_freelancers cl
		INNER JOIN clients c 
			ON cl.client_id = c.id
		INNER JOIN users u 
			ON u.id = c.id
		WHERE freelancer_id = $1`

	var rows *sql.Rows
	var err error

	if search != "" {
		// Validate search input length to prevent abuse
		if len(search) > 100 {
			return nil, fmt.Errorf("search query too long (max 100 characters)")
		}
		query += ` AND (
			LOWER(c.first_name) LIKE LOWER($2) OR 
			LOWER(c.last_name) LIKE LOWER($2) OR 
			LOWER(u.email) LIKE LOWER($2)
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

	var clients []models.ClientFreelancerJoin

	for rows.Next() {
		var c models.ClientFreelancerJoin
		if err := rows.Scan(&c.ClientID, &c.FirstName, &c.LastName, &c.Email,
			&c.CreatedAt, &c.FreelancerID); err != nil {
			return nil, err
		}
		clients = append(clients, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &clients, nil
}
