package glob

import (
	"math/rand"
)

// text contain all text for all labels. Lang of text presets before start of a program or in a settings menu
type text struct {
	StartGame          string   // text for startButton on main page
	Settings           string   // text for settingsButton on main page
	PreGame            string   // pre game label
	HowManyPlayers     string   // text for a label
	HowManyTeams       string   // text for a label
	GameMode           string   // text for a label
	GameModes          []string // array with game modes
	PlayerName         string   // text for a column in a single mode
	ColumnsForTeamMode []string // array with names for a columns in a team mode
	AddThirdPlayer     string   // text for a addThirdPlayerButton in a table
	Language           string   // text for a label
	Languages          []string // array with supported lang
	Difficulty         string   // text for a label
	LevelsOfDifficulty []string // array with difficulties
	Quit               string   // Header for quit menu
	AUSureExit         string   // text for exit a game and go to menu
	AUSureQuit         string   // text to quit a game
	Guesser            string   // text for a label
	Skip               string   // skipButton
	Guessed            string   // guessedButton
	Start              string   // startGameButton
	NextRound          string   // nextRoundButton
	EndGame            string   // text for a button, when game is over
}

// global settings
type settings struct {
	CostOfGuessing int // how much points player will get, when player guesses a word
	CostOfSkip     int // how much points player will lose, when player skips a word
	TimeOfRound    int // how much seconds round goes
	TargetScore    int // after which score game will be over
}

// AllText used to read all text from json file
type AllText struct {
	Ru text
	En text
}

var Text text // All labels text here

var Config settings // Global settings for a game

var PathToUser string // Path to the user dir

// RandIntervalNoRepeat returns array of integers from [0,r-1] without repeat
func RandIntervalNoRepeat(r int) []int {
	m := make([]int, r)
	for i := 1; i < r+1; i++ {
		j := rand.Intn(i)
		m[i-1] = m[j]
		m[j] = i - 1
	}
	return m
}
