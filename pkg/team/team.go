package team

type Team struct {
	Name           string
	Players        []string
	HowManyPlayers int
	Score          int
}

func NewTeam(names []string) *Team {
	var newTeam Team
	newTeam.Name = names[0]
	newTeam.Players = names[1:]
	return &newTeam
}
