ALTER TABLE users
    ADD CONSTRAINT fk_team
        FOREIGN KEY (team_name)
            REFERENCES team(team_name)
            ON DELETE SET NULL
            ON UPDATE CASCADE;

CREATE INDEX IF NOT EXISTS idx_users_team_active ON users (team_name, is_active) WHERE is_active = TRUE;

CREATE INDEX IF NOT EXISTS idx_pr_author ON pull_request (author_id);

CREATE INDEX IF NOT EXISTS idx_reviewer_pr ON pull_request_reviewer (reviewer_id);