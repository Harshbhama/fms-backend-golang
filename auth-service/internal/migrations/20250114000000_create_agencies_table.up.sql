-- +goose Up

CREATE TABLE agencies (
    id INT PRIMARY KEY,
    name VARCHAR(100) NOT NULL CHECK (char_length(name) BETWEEN 1 AND 100),
    email TEXT NOT NULL CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    website TEXT NULL CHECK (website IS NULL OR website ~* '^(https?://)[A-Za-z0-9.-]+(\.[A-Za-z]{2,})(:[0-9]+)?(/.*)?$'),
    description TEXT NOT NULL CHECK (char_length(description) BETWEEN 50 AND 1000),
    location TEXT NOT NULL CHECK (location = ANY(ARRAY[
        'New York', 'San Francisco', 'London', 'Berlin', 'Toronto', 'Sydney'
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
