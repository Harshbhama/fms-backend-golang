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