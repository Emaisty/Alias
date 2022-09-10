package glob

import "math/rand"

type text struct {
	StartGame          string
	Settings           string
	PreGame            string
	HowManyPlayers     string
	HowManyTeams       string
	GameMode           string
	GameModes          []string
	PlayerName         string
	ColumnsForTeamMode []string
	AddThirdPlayer     string
	Language           string
	Languages          []string
	Difficulty         string
	LevelsOfDifficulty []string
	Quit               string
	AUSureExit         string
	AUSureQuit         string
	Guesser            string
	Skip               string
	Guessed            string
	Start              string
	NextRound          string
	EndGame            string
}

type settings struct {
	CostOfGuessing int
	CostOfSkip     int
	TimeOfRound    int
	TargetScore    int
}

type AllText struct {
	Ru text
	En text
}

var Text text

var Config settings

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
