-- Create shortner table
CREATE TABLE IF NOT EXISTS shortner (
    id BIGSERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create index on original_url for faster lookups
CREATE INDEX IF NOT EXISTS idx_shortner_original_url ON shortner(original_url);
