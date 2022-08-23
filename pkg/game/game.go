package game

import (
	"Alias/pkg/team"
	"strconv"
)

type Game struct {
	teams           []team.Team
	whichTurnInTeam []int
	whichTeamTurn   int
}

func NewGame(names [][]string) *Game {
	var game Game
	if len(names[0]) == 1 {
		for _, name1 := range names {
			for _, name2 := range names {
				if name1[0] != name2[0] {
					game.teams = append(game.teams, *team.NewTeam([]string{"", name1[0], name2[0]}))
					game.whichTurnInTeam = append(game.whichTurnInTeam, 0)
				}
			}
		}
	} else {
		for _, oneTeam := range names {
			game.teams = append(game.teams, *team.NewTeam(oneTeam))
			game.whichTurnInTeam = append(game.whichTurnInTeam, 0)
		}
	}
	return &game
}

// GetCurrentPlayersName
// Return current players names.
// If it is TEAM mode - returning (teamName, player1Name, player2Name)
// if it is SOLO mode - returning ("", player1Name, player2Name)
func (game *Game) GetCurrentPlayersName() (string, string, string) {
	currentTeam := game.teams[game.whichTeamTurn]
	return currentTeam.Name,
		currentTeam.Players[game.whichTurnInTeam[game.whichTeamTurn]],
		currentTeam.Players[(game.whichTurnInTeam[game.whichTeamTurn]+1)%len(currentTeam.Players)]
}

// GetTeamsAndTheirResult return array of teams (in format [0] - result, [1] - team name, [2:] - players names)
// In solo format, it will be [0] - score, and [1] - players name
func (game *Game) GetTeamsAndTheirResult() [][]string {
	var teams [][]string
	// if game in solo mode
	if game.teams[0].Name == "" {
		var (
			playersNames = []string{game.teams[0].Players[0], game.teams[0].Players[1]}
			i            = 1
		)
		for i < len(game.teams) && game.teams[i].Players[0] != playersNames[1] {
			playersNames = append(playersNames, game.teams[i].Players[1])
			i++
		}
		for _, name := range playersNames {
			var sum int
			for _, t := range game.teams {
				if t.Players[0] == name || t.Players[1] == name {
					sum += t.Score
				}
			}
			teams = append(teams, []string{strconv.Itoa(sum), name})
		}

	} else {
		for i, team := range game.teams {
			teams = append(teams, []string{strconv.Itoa(team.Score), team.Name})
			for _, player := range team.Players {
				teams[i] = append(teams[i], player)

			}
		}
	}

	return teams
}

// SaveResultAndGoNext add roundScore to team score,
// move guessing role to next player in team (whichTurnInTeam[i]++ % len),
// and moves to next team
func (game *Game) SaveResultAndGoNext(roundScore int) {
	game.teams[game.whichTeamTurn].Score += roundScore
	game.whichTurnInTeam[game.whichTeamTurn] = (game.whichTurnInTeam[game.whichTeamTurn] + 1) % len(game.teams[game.whichTeamTurn].Players)
	game.whichTeamTurn = (game.whichTeamTurn + 1) % len(game.teams)
}

// IsSessionOver answers on question - Had all teams played 1 round?
func (game *Game) IsSessionOver() bool {
	if game.whichTeamTurn == 0 {
		return true
	} else {
		return false
	}
}
