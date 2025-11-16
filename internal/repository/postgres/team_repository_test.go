package postgres

import (
	"AvitoPRService/internal/model/dto"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestTeamRepositoryImpl_CreateTeam(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close() //nolint:errcheck

	repo := NewTeamRepositoryImpl(db)

	teamName := "teamA"
	members := []dto.TeamMemberDto{
		{ID: "1"},
		{ID: "2"},
	}

	// Начало транзакции
	mock.ExpectBegin()
	// Проверка существования команды
	mock.ExpectQuery("SELECT 1 FROM team WHERE team_name = \\$1").
		WithArgs(teamName).
		WillReturnError(sql.ErrNoRows)
	// Вставка команды
	mock.ExpectExec("INSERT INTO team\\(team_name\\) VALUES\\(\\$1\\)").
		WithArgs(teamName).
		WillReturnResult(sqlmock.NewResult(1, 1))
	// Обновление пользователей
	mock.ExpectQuery("UPDATE users SET team_name = \\$1 WHERE id = \\$2 RETURNING id, username, team_name, is_active").
		WithArgs(teamName, "1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "team_name", "is_active"}).AddRow("1", "user1", teamName, true))
	mock.ExpectQuery("UPDATE users SET team_name = \\$1 WHERE id = \\$2 RETURNING id, username, team_name, is_active").
		WithArgs(teamName, "2").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "team_name", "is_active"}).AddRow("2", "user2", teamName, true))
	// Коммит
	mock.ExpectCommit()

	team, err := repo.CreateTeam(teamName, members)
	assert.NoError(t, err)
	assert.Equal(t, teamName, team.Name)
	assert.Len(t, team.Members, 2)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestTeamRepositoryImpl_CreateTeam_AlreadyExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close() //nolint:errcheck

	repo := NewTeamRepositoryImpl(db)
	teamName := "teamA"

	// Метод CreateTeam начинает транзакцию
	mock.ExpectBegin()

	// Проверка существования команды
	mock.ExpectQuery("SELECT 1 FROM team WHERE team_name = \\$1").
		WithArgs(teamName).
		WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

	// Так как команда существует, транзакция не будет коммититься
	mock.ExpectRollback()

	team, err := repo.CreateTeam(teamName, []dto.TeamMemberDto{})
	assert.Nil(t, team)
	assert.ErrorIs(t, err, ErrTeamExists)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestTeamRepositoryImpl_FindTeamByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close() //nolint:errcheck

	repo := NewTeamRepositoryImpl(db)
	teamName := "teamA"

	// Начало транзакции
	mock.ExpectBegin()

	// Проверка существования команды (IsTeamExists)
	mock.ExpectQuery("SELECT 1 FROM team WHERE team_name = \\$1").
		WithArgs(teamName).
		WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

	// Запрос участников команды
	mock.ExpectQuery("SELECT u.id, u.username, u.is_active FROM users AS u JOIN team AS t ON t.team_name = u.team_name WHERE u.team_name = \\$1;").
		WithArgs(teamName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "is_active"}).
			AddRow("1", "user1", true).
			AddRow("2", "user2", true))

	// Так как defer Rollback вызовется всегда
	mock.ExpectRollback()

	team, err := repo.FindTeamByName(teamName)
	assert.NoError(t, err)
	assert.Equal(t, teamName, team.Name)
	assert.Len(t, team.Members, 2)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
