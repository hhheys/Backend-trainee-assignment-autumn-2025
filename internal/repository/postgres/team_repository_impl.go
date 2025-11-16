package postgres

import (
	db2 "AvitoPRService/internal/model/db"
	"AvitoPRService/internal/model/dto"
	"database/sql"
	"errors"
	"log"
)

// TeamRepositoryImpl implements TeamRepository
type TeamRepositoryImpl struct {
	db *sql.DB
}

// NewTeamRepositoryImpl returns new TeamRepositoryImpl
func NewTeamRepositoryImpl(db *sql.DB) *TeamRepositoryImpl {
	return &TeamRepositoryImpl{db: db}
}

// CreateTeam creates a new team
func (r *TeamRepositoryImpl) CreateTeam(teamName string, members []dto.TeamMemberDto) (*db2.Team, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Println("rollback error:", err) // TODO: logger
		}
	}()

	// Check if team exists
	teamExists, err := r.IsTeamExists(teamName)
	if err != nil {
		return nil, err
	}
	if teamExists {
		return nil, ErrTeamExists
	}

	// Team creation
	_, err = tx.Exec("INSERT INTO team(team_name) VALUES($1)", teamName)
	if err != nil {
		return nil, err
	}

	// Link users to the team
	var (
		user        db2.User
		teamMembers = make([]db2.User, 0, len(members))
	)
	for _, member := range members {
		err := tx.QueryRow("UPDATE users SET team_name = $1 WHERE id = $2 RETURNING id, username, team_name, is_active", teamName, member.ID).Scan(&user.ID, &user.Username, &user.TeamName, &user.IsActive)
		if err != nil {
			continue
		}
		teamMembers = append(teamMembers, user)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return db2.NewTeam(teamName, teamMembers), nil
}

// IsTeamExists checks if team exists
func (r *TeamRepositoryImpl) IsTeamExists(teamName string) (bool, error) {
	err := r.db.QueryRow("SELECT 1 FROM team WHERE team_name = $1", teamName).Scan()
	if !errors.Is(err, sql.ErrNoRows) {
		return true, nil
	}
	return false, nil
}

// FindTeamByName finds team by name
func (r *TeamRepositoryImpl) FindTeamByName(teamName string) (*db2.Team, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Println("rollback error:", err) // TODO: logger
		}
	}()

	teamExists, err := r.IsTeamExists(teamName)
	if err != nil {
		return nil, err
	}
	if !teamExists {
		return nil, ErrTeamNotExists
	}

	rows, err := tx.Query(
		"SELECT u.id, u.username, u.is_active FROM users AS u JOIN team AS t ON t.team_name = u.team_name WHERE u.team_name = $1;",
		teamName,
	)
	if err != nil {
		return &db2.Team{}, err
	}
	var teamMembers []db2.User
	var user db2.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.IsActive)
		if err != nil {
			continue
		}
		teamMembers = append(teamMembers, user)
	}
	return db2.NewTeam(teamName, teamMembers), nil
}
