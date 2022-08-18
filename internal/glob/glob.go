package glob

type text struct {
	StartGame      string
	Settings       string
	PreGame        string
	HowManyPlayers string
	HowManyTeams   string
	SoloMode       string
	TeamMode       string
	TeamName       string
	PlayerName     string
	FirstPlayer    string
	SecondPlayer   string
	ThirdPlayer    string
	AddThirdPlayer string
	Language       string
	Russian        string
	English        string
	Difficulty     string
}

type AllText struct {
	Ru text
	En text
}

var Text text
