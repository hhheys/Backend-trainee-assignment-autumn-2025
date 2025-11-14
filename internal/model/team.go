package model

// Team represents a team in the service
type Team struct {
	TeamName string `json:"team_name"`
	Members  []User `json:"members"`
}

// NewTeam creates a new team
func NewTeam(teamName string, members []User) *Team {
	return &Team{
		TeamName: teamName,
		Members:  members,
	}
}
