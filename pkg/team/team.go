package team

type Player struct {
	Name            string
	CompoundedWords int
	GuessedWords    int
}

func NewPlayer(name string) *Player {
	var player Player
	player.Name = name
	return &player
}

type Team struct {
	Name         string
	FirstPlayer  *Player
	SecondPlayer *Player
	ThirdPlayer  *Player
}

func NewTeam(names []string) *Team {
	var newTeam Team
	newTeam.Name = names[0]
	newTeam.FirstPlayer = NewPlayer(names[1])
	newTeam.SecondPlayer = NewPlayer(names[2])
	if names[3] != "" {
		newTeam.ThirdPlayer = NewPlayer(names[3])
	} else {
		newTeam.ThirdPlayer = nil
	}
	return &newTeam
}
