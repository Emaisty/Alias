package game

import (
	"Alias/pkg/team"
)

type Game struct {
	Teams           []team.Team
	WhichTurnInTeam []int
	WhichTeamTurn   int
}

func NewGame(names [][]string) *Game {
	var game Game
	if len(names[0]) == 1 {
		for _, name1 := range names {
			for _, name2 := range names {
				if name1[0] != name2[0] {
					game.Teams = append(game.Teams, *team.NewTeam([]string{"", name1[0], name2[0]}))
					game.WhichTurnInTeam = append(game.WhichTurnInTeam, 0)
				}
			}
		}
	} else {
		for _, oneTeam := range names {
			game.Teams = append(game.Teams, *team.NewTeam(oneTeam))
			game.WhichTurnInTeam = append(game.WhichTurnInTeam, 0)
		}
	}
	return &game
}

// GetCurrentPlayersName
// Return current players names. If it is Team mode - returning (teamName, player1Name, player2Name)
func (game *Game) GetCurrentPlayersName() (string, string, string) {
	currentTeam := game.Teams[game.WhichTeamTurn]
	return currentTeam.Name,
		currentTeam.Players[game.WhichTurnInTeam[game.WhichTeamTurn]],
		currentTeam.Players[(game.WhichTurnInTeam[game.WhichTeamTurn]+1)%len(currentTeam.Players)]
}

func (game *Game) SaveResultAndGoNext(res int) {
	game.Teams[game.WhichTeamTurn].Score += res
	game.WhichTurnInTeam[game.WhichTeamTurn] = (game.WhichTurnInTeam[game.WhichTeamTurn] + 1) % len(game.Teams[game.WhichTeamTurn].Players)
	game.WhichTeamTurn = (game.WhichTeamTurn + 1) % len(game.Teams)
}

func (game *Game) IsSessionOver() bool {
	if game.WhichTeamTurn == 0 {
		return true
	} else {
		return false
	}
}
