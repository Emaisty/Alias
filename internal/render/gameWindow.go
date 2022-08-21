package render

import (
	"Alias/internal/glob"
	"Alias/pkg/game"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"math/rand"
	"strconv"
)

type gamePageObjects struct {
	application
	language            string
	forceQuitButton     *widgets.QPushButton
	timer               *core.QTimer
	downCounter         int
	words               []string
	whichWordChoose     []int
	currentWord         int
	guessedWords        []string
	HadBeenGuessed      []bool
	timerBoard          *widgets.QLCDNumber
	wordLabel           *widgets.QLabel
	teamNameLabel       *widgets.QLabel
	firstPlayerLabel    *widgets.QLabel
	secondPlayerLabel   *widgets.QLabel
	skipButton          *widgets.QPushButton
	nextWordButton      *widgets.QPushButton
	startRoundButton    *widgets.QPushButton
	goToNextRoundButton *widgets.QPushButton
	game                *game.Game
}

func newGamePage(app application, names [][]string, lang int) *gamePageObjects {
	var gamePage gamePageObjects
	gamePage.application = app
	gamePage.game = game.NewGame(names)
	if lang == 0 {
		gamePage.language = "rus"
	} else {
		gamePage.language = "eng"
	}
	return &gamePage
}

func (page *gamePageObjects) createLabels() {
	page.wordLabel = widgets.NewQLabel2("Word", nil, 0)
	page.teamNameLabel = widgets.NewQLabel2("Team Name", nil, 0)
	page.firstPlayerLabel = widgets.NewQLabel2("First Player Name", nil, 0)
	page.secondPlayerLabel = widgets.NewQLabel2("Second Player Name", nil, 0)
}

func (page *gamePageObjects) quitGame() {
	res := widgets.QMessageBox_Question(nil, glob.Text.Quit, glob.Text.Quit, widgets.QMessageBox__Close|widgets.QMessageBox__No|widgets.QMessageBox__Yes, widgets.QMessageBox__Yes)
	if res == widgets.QMessageBox__Yes {
		page.timer.Stop()
		page.window.DisconnectCloseEvent()
		page.application.displayMainMenu()
	}
}

func (page *gamePageObjects) startRound() {
	page.timer.Start2()
	page.roundPageRender()
}

func (page *gamePageObjects) goNextRound() {
	var resultScore int
	for _, hadBennGuessed := range page.HadBeenGuessed {
		if hadBennGuessed {
			resultScore += 1
		} else {
			resultScore -= glob.Config.CostOfSkip
		}
	}
	page.game.SaveResultAndGoNext(resultScore)

	if page.game.IsSessionOver() {
		page.resultPageRender()
	} else {
		page.render()
	}

}

func (page *gamePageObjects) wordHadBeenGuessed() {
	page.guessedWords = append(page.guessedWords, page.words[page.currentWord])
	page.HadBeenGuessed = append(page.HadBeenGuessed, true)

	page.currentWord = (page.currentWord + 1) % len(page.words)

	page.wordLabel.SetText(page.words[page.whichWordChoose[page.currentWord]])
}

func (page *gamePageObjects) wordHadNotBeenGuessed() {
	page.guessedWords = append(page.guessedWords, page.words[page.currentWord])
	page.HadBeenGuessed = append(page.HadBeenGuessed, false)

	page.currentWord = (page.currentWord + 1) % len(page.words)

	page.wordLabel.SetText(page.words[page.whichWordChoose[page.currentWord]])
}

func (page *gamePageObjects) createButtons() {
	page.forceQuitButton = widgets.NewQPushButton2("<", nil)
	page.forceQuitButton.ConnectPressed(page.quitGame)

	page.skipButton = widgets.NewQPushButton2(glob.Text.Skip, nil)
	page.skipButton.ConnectPressed(page.wordHadNotBeenGuessed)

	page.nextWordButton = widgets.NewQPushButton2(glob.Text.Guessed, nil)
	page.nextWordButton.ConnectPressed(page.wordHadBeenGuessed)

	page.startRoundButton = widgets.NewQPushButton2(glob.Text.Start, nil)
	page.startRoundButton.ConnectPressed(page.startRound)

	page.goToNextRoundButton = widgets.NewQPushButton2(glob.Text.NextRound, nil)
	page.goToNextRoundButton.ConnectPressed(page.goNextRound)
}

func (page *gamePageObjects) onTimer() {
	page.downCounter--

	min := page.downCounter / 60
	sec := page.downCounter % 60

	page.timerBoard.Display(strconv.Itoa(min) + ":" + strconv.Itoa(sec))

	if page.downCounter == 0 {
		page.timer.Stop()
		page.postRoundPageRender()
	}
}

func (page *gamePageObjects) createTimer() {
	page.timer = core.NewQTimer(nil)
	page.timer.SetInterval(1000)
	page.timer.ConnectTimeout(page.onTimer)
}

func (page *gamePageObjects) createObjects() {
	page.createLabels()

	page.timerBoard = widgets.NewQLCDNumber(nil)

	page.createTimer()

	page.createButtons()

}

func (page *gamePageObjects) listOfWords(layout *widgets.QGridLayout) {
	var resultScore int
	for _, hadBennGuessed := range page.HadBeenGuessed {
		if hadBennGuessed {
			resultScore += 1
		} else {
			resultScore -= glob.Config.CostOfSkip
		}
	}

	resLabel := widgets.NewQLabel2(strconv.Itoa(resultScore), nil, 0)

	layout.AddWidget(resLabel)

	for i, word := range page.guessedWords {
		wordLabel := widgets.NewQLabel2(word, nil, 0)
		isWordGuessed := widgets.NewQCheckBox(nil)
		isWordGuessed.SetChecked(page.HadBeenGuessed[i])
		isWordGuessed.ConnectClicked(func(checked bool) {
			page.HadBeenGuessed[i] = checked
			if checked {
				resultScore += glob.Config.CostOfSkip + 1
			} else {
				resultScore -= glob.Config.CostOfSkip + 1
			}
			resLabel.SetText(strconv.Itoa(resultScore))
		})
		layout.AddWidget(wordLabel)
		layout.AddWidget(isWordGuessed)
	}
}

func (page *gamePageObjects) postRoundPageRender() {
	layout := widgets.NewQGridLayout2()

	layout.AddWidget(page.forceQuitButton)

	page.listOfWords(layout)

	layout.AddWidget(page.goToNextRoundButton)

	page.application.show(layout)
}

func (page *gamePageObjects) resultPageRender() {
	layout := widgets.NewQGridLayout2()

	layout.AddWidget(page.forceQuitButton)

	targetLabel := widgets.NewQLabel2("Target is "+strconv.Itoa(glob.Config.TargetScore), nil, 0)

	layout.AddWidget(targetLabel)

	var maxScore int

	for _, team := range page.game.Teams {
		layout.AddWidget(widgets.NewQLabel2(team.Name, nil, 0))
		layout.AddWidget(widgets.NewQLabel2(team.Players[0], nil, 0))
		layout.AddWidget(widgets.NewQLabel2(team.Players[1], nil, 0))
		if len(team.Players) == 3 {
			layout.AddWidget(widgets.NewQLabel2(team.Players[2], nil, 0))
		}
		if team.Score > maxScore {
			maxScore = team.Score
		}
		layout.AddWidget(widgets.NewQLabel2(strconv.Itoa(team.Score), nil, 0))
	}
	nextActionButton := widgets.NewQPushButton(nil)
	if maxScore > glob.Config.TargetScore {
		nextActionButton.SetText(glob.Text.EndGame)
		nextActionButton.ConnectPressed(func() {
			page.application.window.DisconnectCloseEvent()
			page.application.displayMainMenu()
		})
	} else {
		nextActionButton.SetText(glob.Text.NextRound)
		nextActionButton.ConnectPressed(page.render)
	}
	layout.AddWidget(nextActionButton)

	page.application.show(layout)
}

func (page *gamePageObjects) roundPageRender() {
	layout := widgets.NewQGridLayout2()

	layout.AddWidget(page.forceQuitButton)
	layout.AddWidget(page.timerBoard)
	layout.AddWidget(page.wordLabel)
	layout.AddWidget(page.skipButton)
	layout.AddWidget(page.nextWordButton)

	page.application.show(layout)
}

func (page *gamePageObjects) prepareTimerForNewRound() {
	page.downCounter = glob.Config.TimeOfRound

	min := page.downCounter / 60
	sec := page.downCounter % 60

	page.timerBoard.Display(strconv.Itoa(min) + ":" + strconv.Itoa(sec))

}

func randInterval_norepeat(r int) []int {
	m := make([]int, r)
	for i := 1; i < r+1; i++ {
		j := rand.Intn(i)
		m[i-1] = m[j]
		m[j] = i - 1
	}
	return m
}

func getWordsFromTable(lang string) ([]string, []int) {
	db, err := sql.Open("sqlite3", "data/words")
	if err != nil {
		panic(err.Error())

	}
	query, errr := db.Query("SELECT word from words WHERE language = '" + lang + "'")
	if errr != nil {
		panic(err.Error())
	}
	var str string
	var res []string
	for query.Next() {
		query.Scan(&str)
		res = append(res, str)
	}
	return res, randInterval_norepeat(len(res))
}

func (page *gamePageObjects) prepareWordsForNewRound() {
	page.words, page.whichWordChoose = getWordsFromTable(page.language)
	page.currentWord = 0

	page.guessedWords = make([]string, 0)
	page.HadBeenGuessed = make([]bool, 0)

	page.wordLabel.SetText(page.words[page.whichWordChoose[page.currentWord]])

}

func (page *gamePageObjects) preRoundPageRender() {
	layout := widgets.NewQGridLayout2()

	teamName, firstPlayerName, SecondPlayerName := page.game.GetCurrentPlayersName()

	// Prepare timer for new round
	page.prepareTimerForNewRound()

	page.prepareWordsForNewRound()

	page.teamNameLabel.SetText(teamName)
	page.firstPlayerLabel.SetText(firstPlayerName)
	page.secondPlayerLabel.SetText(SecondPlayerName)

	layout.AddWidget(page.forceQuitButton)
	layout.AddWidget(page.teamNameLabel)
	layout.AddWidget(page.firstPlayerLabel)
	layout.AddWidget(page.secondPlayerLabel)
	layout.AddWidget(page.startRoundButton)

	page.application.show(layout)
}

func (page *gamePageObjects) render() {
	page.createObjects()

	page.preRoundPageRender()
}

// parent - window of app; teamName - raw array with names. NewGame([][]string) will work with it;
// language: 0 - Ru, 1 - En
func (app application) displayGameWindow(teamsName [][]string, language int) {
	// Connect close event, to prevent random quit
	app.window.ConnectCloseEvent(closeEvent)

	gamePage := newGamePage(app, teamsName, language)

	gamePage.render()
}
