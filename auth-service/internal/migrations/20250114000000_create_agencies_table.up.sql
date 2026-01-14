-- +goose Up

CREATE TABLE agencies (
    id BIGINT PRIMARY KEY,
    name TEXT NULL,
    email TEXT NULL,
    website TEXT NULL,
    description TEXT NULL,
    location TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_agency_user FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE
);
