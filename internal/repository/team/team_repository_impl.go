package team

import (
	"AvitoPRService/internal/db"
	"AvitoPRService/internal/dto"
	"AvitoPRService/internal/model"
	"database/sql"
	"errors"
)

type TeamRepositoryImpl struct {
	db *sql.DB
}

func NewTeamRepositoryImpl(db *sql.DB) *TeamRepositoryImpl {
	return &TeamRepositoryImpl{db: db}
}

func (r *TeamRepositoryImpl) CreateTeam(teamName string, members []dto.TeamMemberDto) (*model.Team, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	// Check if team exists
	teamExists, err := r.IsTeamExists(teamName)
	if err != nil {
		return nil, err
	}
	if teamExists {
		return nil, db.ErrTeamExists
	}

	// Team creation
	_, err = tx.Exec("INSERT INTO team(team_name) VALUES($1)", teamName)
	if err != nil {
		return nil, err
	}

	// Link users to the team
	var (
		user        model.User
		teamMembers []model.User = make([]model.User, 0, len(members))
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
	} else {
		return model.NewTeam(teamName, teamMembers), nil
	}

}

func (r *TeamRepositoryImpl) IsTeamExists(teamName string) (bool, error) {
	err := r.db.QueryRow("SELECT 1 FROM team WHERE team_name = $1", teamName).Scan()
	if !errors.Is(err, sql.ErrNoRows) {
		return true, nil
	}
	return false, nil
}

func (r *TeamRepositoryImpl) FindTeamByName(teamName string) (*model.Team, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	teamExists, err := r.IsTeamExists(teamName)
	if err != nil {
		return nil, err
	}
	if !teamExists {
		return nil, db.ErrTeamNotExists
	}

	rows, err := tx.Query(
		"SELECT u.id, u.username, u.is_active FROM users AS u JOIN team AS t ON t.team_name = u.team_name WHERE u.team_name = $1;",
		teamName,
	)
	if err != nil {
		return &model.Team{}, err
	}
	var teamMembers []model.User
	var user model.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.IsActive)
		if err != nil {
			continue
		}
		teamMembers = append(teamMembers, user)
	}
	return model.NewTeam(teamName, teamMembers), nil
}
