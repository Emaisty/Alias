package glob

type text struct {
	StartGame          string
	Settings           string
	PreGame            string
	HowManyPlayers     string
	HowManyTeams       string
	GameMode           string
	SoloMode           string
	TeamMode           string
	PlayerName         string
	ColumnsForTeamMode []string
	AddThirdPlayer     string
	Language           string
	Russian            string
	English            string
	Difficulty         string
	Quit               string
	AUSureExit         string
	AUSureQuit         string
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
