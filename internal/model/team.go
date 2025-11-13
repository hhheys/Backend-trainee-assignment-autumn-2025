package model

type Team struct {
	TeamName string `json:"team_name"`
	Members  []User `json:"members"`
}

func NewTeam(teamName string, members []User) *Team {
	return &Team{
		TeamName: teamName,
		Members:  members,
	}
}
