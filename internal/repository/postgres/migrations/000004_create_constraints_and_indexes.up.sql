ALTER TABLE users
    ADD CONSTRAINT fk_team
        FOREIGN KEY (team_name)
            REFERENCES team(team_name)
            ON DELETE SET NULL
            ON UPDATE CASCADE;

