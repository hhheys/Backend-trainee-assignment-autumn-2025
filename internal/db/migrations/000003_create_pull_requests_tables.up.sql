CREATE TYPE pull_request_status AS ENUM ('OPEN', 'MERGED');

CREATE TABLE IF NOT EXISTS pull_request (
    id VARCHAR(255) PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    author_id UUID NOT NULL,
    status pull_request_status NOT NULL DEFAULT 'OPEN',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    merged_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT pull_request_author_id_fk
    FOREIGN KEY(author_id)
    REFERENCES users(id)
    ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS pull_request_reviewer (
    pull_request_id VARCHAR(255) NOT NULL,
    reviewer_id UUID NOT NULL,

    PRIMARY KEY (pull_request_id, reviewer_id),

    assigned_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT pull_request_reviewer_pull_request_id_fk
    FOREIGN KEY(pull_request_id)
    REFERENCES pull_request(id)
    ON DELETE CASCADE,

    CONSTRAINT pull_request_reviewer_user_id_fk
    FOREIGN KEY(reviewer_id)
    REFERENCES users(id)
    ON DELETE RESTRICT
);