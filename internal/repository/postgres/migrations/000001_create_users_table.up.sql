CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT (gen_random_uuid()),
    username VARCHAR(255) NOT NULL,
    team_name VARCHAR(255),
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);
