package game

import "Alias/pkg/team"

type Game struct {
	Players []team.Player
	Teams   []team.Team
}

func NewGame(names [][]string) *Game {
	var game Game
	if len(names[0]) == 1 {
		game.Teams = nil
		for _, name := range names {
			game.Players = append(game.Players, *team.NewPlayer(name[0]))
		}
	} else {
		game.Players = nil
		for _, teams := range names {
			game.Teams = append(game.Teams, *team.NewTeam(teams))
		}
	}
	return &game
}
