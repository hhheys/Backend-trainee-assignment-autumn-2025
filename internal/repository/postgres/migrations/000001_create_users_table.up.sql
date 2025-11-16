CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT (gen_random_uuid()),
    username VARCHAR(255) NOT NULL,
    team_name VARCHAR(255),
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE INDEX IF NOT EXISTS idx_users_team_active ON users (team_name, is_active) WHERE is_active = TRUE;

CREATE INDEX IF NOT EXISTS idx_pr_author ON pull_request (author_id);

CREATE INDEX IF NOT EXISTS idx_reviewer_pr ON pull_request_reviewer (reviewer_id);
