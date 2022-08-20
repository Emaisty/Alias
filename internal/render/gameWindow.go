package render

import (
	"Alias/internal/glob"
	"Alias/pkg/game"
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"strconv"
)

type gamePageObjects struct {
	application
	forceQuitButton     *widgets.QPushButton
	timer               *core.QTimer
	downCounter         int
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

func newGamePage(app application, names [][]string) *gamePageObjects {
	var gamePage gamePageObjects
	gamePage.application = app
	gamePage.game = game.NewGame(names)
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
		page.application.displayMainMenu()
	}
}

func (page *gamePageObjects) startRound() {
	page.timer.Start2()
	page.roundPageRender()
}

func (page *gamePageObjects) goNextRound() {
	//TODO save results from round
	page.preRoundPageRender()
}

func (page *gamePageObjects) createButtons() {
	page.forceQuitButton = widgets.NewQPushButton2("<", nil)
	page.forceQuitButton.ConnectPressed(page.quitGame)

	page.skipButton = widgets.NewQPushButton2(glob.Text.Skip, nil)
	page.skipButton.ConnectPressed(func() {

	})

	page.nextWordButton = widgets.NewQPushButton2(glob.Text.Guessed, nil)
	page.nextWordButton.ConnectPressed(func() {

	})

	page.startRoundButton = widgets.NewQPushButton2(glob.Text.Start, nil)
	page.startRoundButton.ConnectPressed(page.startRound)

	page.goToNextRoundButton = widgets.NewQPushButton2(glob.Text.NextRound, nil)
	page.goToNextRoundButton.ConnectPressed(func() {

	})
}

func (page *gamePageObjects) onTimer() {
	fmt.Println(page.downCounter)
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

func (page *gamePageObjects) postRoundPageRender() {
	layout := widgets.NewQGridLayout2()

	layout.AddWidget(page.forceQuitButton)
	//TODO add here words, which had been guessed
	layout.AddWidget(page.goToNextRoundButton)

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

func (page *gamePageObjects) preRoundPageRender() {
	layout := widgets.NewQGridLayout2()

	teamName, firstPlayerName, SecondPlayerName := page.game.GetCurrentPlayersName()

	// Prepare timer for new round
	page.prepareTimerForNewRound()

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
	page.preRoundPageRender()
}

// parent - window of app; teamName - raw array with names. NewGame([][]string) will work with it;
// language: 0 - Ru, 1 - En
func (app application) displayGameWindow(teamsName [][]string, language int) {
	// Connect close event, to prevent random quit
	app.window.ConnectCloseEvent(closeEvent)

	gamePage := newGamePage(app, teamsName)

	gamePage.createObjects()

	gamePage.render()
}
