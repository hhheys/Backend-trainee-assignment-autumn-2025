package db

// Team represents a team in the service
type Team struct {
	Name    string `json:"name"`
	Members []User `json:"members"`
}

// NewTeam creates a new team
func NewTeam(teamName string, members []User) *Team {
	return &Team{
		Name:    teamName,
		Members: members,
	}
}
