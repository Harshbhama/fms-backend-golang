-- +goose Up

CREATE TABLE clients (
	id BIGINT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_client_user FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE
);


ALTER TABLE clients DROP COLUMN name, DROP COLUMN email;
ALTER TABLE clients ADD COLUMN first_name VARCHAR(255) NOT NULL;
ALTER TABLE clients ADD COLUMN last_name VARCHAR(255) NOT NULL;

-- +goose Down
DROP TABLE IF EXISTS clients;





CREATE TABLE freelancers (
    id INT PRIMARY KEY,
    first_name TEXT NULL,
    last_name TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_client_user FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE
);




CREATE TABLE client_freelancers (
    client_id BIGINT NOT NULL,
    freelancer_id INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
    CONSTRAINT fk_freelancer FOREIGN KEY (freelancer_id) REFERENCES freelancers(id) ON DELETE CASCADE,
    
    CONSTRAINT pk_client_freelancer PRIMARY KEY (client_id, freelancer_id)
);


CREATE TABLE client_agencies (
    client_id BIGINT NOT NULL,
    agency_id INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_client_agency FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
    CONSTRAINT fk_agency FOREIGN KEY (agency_id) REFERENCES agency(id) ON DELETE CASCADE,
    
    CONSTRAINT pk_client_agency PRIMARY KEY (client_id, agency_id)
);


CREATE TABLE freelancer_rates (
    id BIGSERIAL PRIMARY KEY,
    freelancer_id INT NOT NULL,
    rate NUMERIC(10, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL, -- ISO 4217 format (e.g., USD, INR, EUR)
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_freelancer FOREIGN KEY (freelancer_id) REFERENCES freelancers(id) ON DELETE CASCADE
);


ALTER TABLE freelancer_rates
ADD CONSTRAINT uq_freelancer_id UNIQUE (freelancer_id);




CREATE TABLE projects (	
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    status VARCHAR(50) NOT NULL, -- e.g., 'pending', 'in_progress', 'completed'
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE client_projects (
    client_id BIGINT NOT NULL,
    project_id INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
    CONSTRAINT fk_project FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    
    CONSTRAINT pk_client_project PRIMARY KEY (client_id, project_id)
);


CREATE TABLE freelancer_projects (
    freelancer_id BIGINT NOT NULL,
    project_id INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_freelancer FOREIGN KEY (freelancer_id) REFERENCES freelancers(id) ON DELETE CASCADE,
    CONSTRAINT fk_project FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    
    CONSTRAINT pk_freelancer_project PRIMARY KEY (freelancer_id, project_id)
);



CREATE TABLE freelancer_timesheet (
    id SERIAL PRIMARY KEY,
    freelancer_id INT NOT NULL,
    project_id INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_timesheet_freelancer FOREIGN KEY (freelancer_id) REFERENCES freelancers(id) ON DELETE CASCADE,
    CONSTRAINT fk_timesheet_project FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);



-- Shard 1
CREATE TABLE timesheetmetadata1 (
    metadata_id SERIAL PRIMARY KEY,
    timesheet_id INT NOT NULL,
    date DATE NOT NULL,
    hours NUMERIC(5,2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    remarks TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_timesheetmetadata1 FOREIGN KEY (timesheet_id) REFERENCES freelancer_timesheet(id) ON DELETE CASCADE
);

-- Shard 2
CREATE TABLE timesheetmetadata2 (
    metadata_id SERIAL PRIMARY KEY,
    timesheet_id INT NOT NULL,
    date DATE NOT NULL,
    hours NUMERIC(5,2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    remarks TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_timesheetmetadata2 FOREIGN KEY (timesheet_id) REFERENCES freelancer_timesheet(id) ON DELETE CASCADE
);

-- Shard 3
CREATE TABLE timesheetmetadata3 (
    metadata_id SERIAL PRIMARY KEY,
    timesheet_id INT NOT NULL,
    date DATE NOT NULL,
    hours NUMERIC(5,2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    remarks TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_timesheetmetadata3 FOREIGN KEY (timesheet_id) REFERENCES freelancer_timesheet(id) ON DELETE CASCADE
);


CREATE TABLE agency (
    id INT PRIMARY KEY,
    name VARCHAR(100) NOT NULL CHECK (char_length(name) BETWEEN 1 AND 100),
    email TEXT NOT NULL CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    website TEXT NULL CHECK (website IS NULL OR website ~* '^(https?://)[A-Za-z0-9.-]+(\.[A-Za-z]{2,})(:[0-9]+)?(/.*)?$'),
    description TEXT NOT NULL CHECK (char_length(description) BETWEEN 50 AND 1000),
    location TEXT NOT NULL CHECK (location = ANY(ARRAY[
        'New York', 'San Francisco', 'London', 'Berlin', 'Toronto', 'Sydney'  -- <-- replace with your allowed locations
    ])),
    team_size INT NOT NULL CHECK (team_size >= 1),
    founded_year INT NOT NULL CHECK (founded_year BETWEEN 1900 AND EXTRACT(YEAR FROM CURRENT_DATE)::INT),
    min_budget NUMERIC(12,2) NOT NULL CHECK (min_budget >= 1000),
    avg_hourly_rate NUMERIC(10,2) NOT NULL CHECK (avg_hourly_rate >= 10),
    specializations TEXT[] NOT NULL CHECK (array_length(specializations, 1) >= 1),
    services TEXT[] NOT NULL CHECK (array_length(services, 1) >= 1),
    phone TEXT NULL,
    address TEXT NULL,
    certifications TEXT[] NULL,
    languages TEXT[] NULL,
    send_invitation BOOLEAN NOT NULL DEFAULT TRUE,
    add_to_favorites BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_agency_user FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE
);

-- Suppose it returns: agency_location_check
ALTER TABLE agency
  DROP CONSTRAINT agency_location_check;


ALTER TABLE projects
ADD COLUMN category TEXT,
ADD COLUMN priority TEXT,
ADD COLUMN required_skills TEXT[],
ADD COLUMN custom_skills TEXT[],
ADD COLUMN detailed_requirements TEXT,
ADD COLUMN expected_deliverables TEXT,
ADD COLUMN assignment_timing TEXT,
ADD COLUMN timeline JSONB;