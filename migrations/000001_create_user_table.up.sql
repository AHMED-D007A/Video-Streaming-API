CREATE TABLE users (
    id SERIAL PRIMARY KEY,               -- Auto-incrementing unique ID for each user
    username VARCHAR(50) NOT NULL, -- Unique username for each user
    email VARCHAR(100) UNIQUE NOT NULL,  -- Unique email for each user
    password_hash TEXT NOT NULL,         -- Hashed password (use bcrypt or similar for hashing)
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, -- Account creation time
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP -- Last update time
);

-- Add a trigger to update the `updated_at` field on any change
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();