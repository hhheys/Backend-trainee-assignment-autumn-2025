CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    team_name VARCHAR(255),
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);