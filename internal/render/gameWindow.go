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

//================================================================================================================
// Class block

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

//================================================================================================================

// randIntervalNoRepeat returns array of integers from [0,r-1] without repeat
func randIntervalNoRepeat(r int) []int {
	m := make([]int, r)
	for i := 1; i < r+1; i++ {
		j := rand.Intn(i)
		m[i-1] = m[j]
		m[j] = i - 1
	}
	return m
}

// getWordsFromTable . Open db and get from it array of words, which on gamePageObjects.language language
func getWordsFromTable(lang string) ([]string, []int) {
	db, _ := sql.Open("sqlite3", "data/words")
	query, _ := db.Query("SELECT word from words WHERE language = '" + lang + "'")
	var str string
	var res []string
	for query.Next() {
		query.Scan(&str)
		res = append(res, str)
	}
	return res, randIntervalNoRepeat(len(res))
}

// Calculate score, which had been received for round
func (page *gamePageObjects) calculateRoundScore() int {
	var resultScore int
	for _, wordHadBeenGuessed := range page.HadBeenGuessed {
		if wordHadBeenGuessed {
			resultScore += glob.Config.CostOfGuessing
		} else {
			resultScore -= glob.Config.CostOfSkip
		}
	}
	return resultScore
}

// exit the game to main menu
func (page *gamePageObjects) exit() {
	page.timer.Stop()
	page.window.DisconnectCloseEvent()
	page.application.displayMainMenuWindow()
}

// Display time on Board in format MM:SS
func (page *gamePageObjects) displayTimeOnBoard() {
	min := page.downCounter / 60
	sec := page.downCounter % 60

	// Display time on Board in format MM:SS
	var (
		minString = strconv.Itoa(min)
		secString = strconv.Itoa(sec)
	)
	if min < 9 {
		minString = "0" + minString
	}
	if sec < 9 {
		secString = "0" + secString
	}
	page.timerBoard.Display(minString + ":" + secString)
}

//================================================================================================================
// Events

func (page *gamePageObjects) backToMeinMenuEvent() {
	res := widgets.QMessageBox_Question(nil, glob.Text.Quit, glob.Text.Quit, widgets.QMessageBox__Close|widgets.QMessageBox__No|widgets.QMessageBox__Yes, widgets.QMessageBox__Yes)
	if res == widgets.QMessageBox__Yes {
		page.exit()
	}
}

// onTimer calls every tick of the timer (every 1 sec)
func (page *gamePageObjects) onTimer() {
	// reduce timer by 1 sec
	page.downCounter--

	page.displayTimeOnBoard()

	if page.downCounter == 0 {
		page.timer.Stop()
		page.postRoundPageRender()
	}
}

// skip the word while the round
func (page *gamePageObjects) wordHadNotBeenGuessed() {
	page.guessedWords = append(page.guessedWords, page.words[page.currentWord])
	page.HadBeenGuessed = append(page.HadBeenGuessed, false)

	page.currentWord = (page.currentWord + 1) % len(page.words)

	page.wordLabel.SetText(page.words[page.whichWordChoose[page.currentWord]])
}

// guessing the word while the round
func (page *gamePageObjects) wordHadBeenGuessed() {
	page.guessedWords = append(page.guessedWords, page.words[page.currentWord])
	page.HadBeenGuessed = append(page.HadBeenGuessed, true)

	page.currentWord = (page.currentWord + 1) % len(page.words)

	page.wordLabel.SetText(page.words[page.whichWordChoose[page.currentWord]])
}

// start round from pre round page
func (page *gamePageObjects) startRound() {
	page.timer.Start2()
	page.roundPageRender()
}

// go to next page (result page or a pre round page)
func (page *gamePageObjects) goAfterPostGamePage() {
	resScore := page.calculateRoundScore()

	page.game.SaveResultAndGoNext(resScore)

	if page.game.IsSessionOver() {
		page.resultPageRender()
	} else {
		page.render()
	}
}

//================================================================================================================
// Create widgets

func (page *gamePageObjects) createLabels() {
	page.wordLabel = widgets.NewQLabel2("Word", nil, 0)
	page.teamNameLabel = widgets.NewQLabel2("Team Name", nil, 0)
	page.firstPlayerLabel = widgets.NewQLabel2("First Player Name", nil, 0)
	page.secondPlayerLabel = widgets.NewQLabel2("Second Player Name", nil, 0)
}

func (page *gamePageObjects) createTimer() {
	page.timer = core.NewQTimer(nil)
	// set tick once per sec¬
	page.timer.SetInterval(1000)
	page.timer.ConnectTimeout(page.onTimer)
}

func (page *gamePageObjects) createButtons() {
	// leave the game to main menu
	page.forceQuitButton = widgets.NewQPushButton2("<", nil)
	page.forceQuitButton.ConnectPressed(page.backToMeinMenuEvent)

	// skip the word while the round
	page.skipButton = widgets.NewQPushButton2(glob.Text.Skip, nil)
	page.skipButton.ConnectPressed(page.wordHadNotBeenGuessed)

	// guessing the word while the round
	page.nextWordButton = widgets.NewQPushButton2(glob.Text.Guessed, nil)
	page.nextWordButton.ConnectPressed(page.wordHadBeenGuessed)

	// start round from pre round page
	page.startRoundButton = widgets.NewQPushButton2(glob.Text.Start, nil)
	page.startRoundButton.ConnectPressed(page.startRound)

	// go to next page (result page or a pre round page)
	page.goToNextRoundButton = widgets.NewQPushButton2(glob.Text.NextRound, nil)
	page.goToNextRoundButton.ConnectPressed(page.goAfterPostGamePage)
}

func (page *gamePageObjects) createObjects() {
	page.createLabels()

	page.timerBoard = widgets.NewQLCDNumber(nil)

	page.createTimer()

	page.createButtons()
}

func (page *gamePageObjects) createLabelsWithTeamsAndTheirResults(layout *widgets.QGridLayout) {
	var maxScore int

	// for each team display their team and players name
	for _, team := range page.game.Teams {
		layout.AddWidget(widgets.NewQLabel2(team.Name, nil, 0))
		layout.AddWidget(widgets.NewQLabel2(team.Players[0], nil, 0))
		layout.AddWidget(widgets.NewQLabel2(team.Players[1], nil, 0))
		if len(team.Players) == 3 {
			layout.AddWidget(widgets.NewQLabel2(team.Players[2], nil, 0))
		}
		layout.AddWidget(widgets.NewQLabel2(strconv.Itoa(team.Score), nil, 0))

		// find max score
		if team.Score > maxScore {
			maxScore = team.Score
		}
	}
	nextActionButton := widgets.NewQPushButton(nil)

	// If score of any team MORE than target score - set on button action to end the game
	// if score LESS - continue the game, go to next session
	if maxScore > glob.Config.TargetScore {
		nextActionButton.SetText(glob.Text.EndGame)
		nextActionButton.ConnectPressed(page.exit)
	} else {
		nextActionButton.SetText(glob.Text.NextRound)
		nextActionButton.ConnectPressed(page.render)
	}

	layout.AddWidget(nextActionButton)
}

//================================================================================================================
// Render

// resultPageRender calls after each session (all team played 1 round)
// Displays all teams with players names and their score
func (page *gamePageObjects) resultPageRender() {
	// Label, which display required result to end the match
	targetLabel := widgets.NewQLabel2("Target is "+strconv.Itoa(glob.Config.TargetScore), nil, 0)

	layout := widgets.NewQGridLayout2()

	layout.AddWidget(page.forceQuitButton)

	layout.AddWidget(targetLabel)

	page.createLabelsWithTeamsAndTheirResults(layout)

	page.application.show(layout)
}

// renderListOfWords display words on post round page
// Status of word (guessed/skipped) can be changed by using check box
func (page *gamePageObjects) renderListOfWords(layout *widgets.QGridLayout) {

	resScore := page.calculateRoundScore()

	resLabel := widgets.NewQLabel2(strconv.Itoa(resScore), nil, 0)

	layout.AddWidget(resLabel)

	// Display all words and check box near them
	// Check box: checked - word had been guessed. not checked - had been skipped
	for i, word := range page.guessedWords {
		// Label with word
		wordLabel := widgets.NewQLabel2(word, nil, 0)
		// Create check box
		isWordGuessed := widgets.NewQCheckBox(nil)
		// If word had been guessed - set checked
		isWordGuessed.SetChecked(page.HadBeenGuessed[i])
		// Connect change of check box. Checked - word had been guessed. Not - skipped
		// Recalculate result score
		isWordGuessed.ConnectClicked(func(checked bool) {
			page.HadBeenGuessed[i] = checked
			if checked {
				resScore += glob.Config.CostOfSkip + glob.Config.CostOfGuessing
			} else {
				resScore -= glob.Config.CostOfSkip + glob.Config.CostOfGuessing
			}
			resLabel.SetText(strconv.Itoa(resScore))
		})

		layout.AddWidget(wordLabel)
		layout.AddWidget(isWordGuessed)
	}
}

// Render post round page. Display score and words. Can rechoose, which had been guessed, and which not
func (page *gamePageObjects) postRoundPageRender() {
	layout := widgets.NewQGridLayout2()

	layout.AddWidget(page.forceQuitButton)

	page.renderListOfWords(layout)

	layout.AddWidget(page.goToNextRoundButton)

	page.application.show(layout)
}

// Set on timer start time
func (page *gamePageObjects) prepareTimerForNewRound() {
	page.downCounter = glob.Config.TimeOfRound

	// Display time on Board in format MM:SS
	page.displayTimeOnBoard()

}

// Load words from db and shuffle them
func (page *gamePageObjects) prepareWordsForNewRound() {
	// load words from db and random shuffle for words
	page.words, page.whichWordChoose = getWordsFromTable(page.language)
	page.currentWord = 0

	// clear array for guessed words
	page.guessedWords = make([]string, 0)
	page.HadBeenGuessed = make([]bool, 0)

	// set 1st word to label
	page.wordLabel.SetText(page.words[page.whichWordChoose[page.currentWord]])

}

// Render round page. Display timer, word and 2 buttons: skip and next word.
// When timer will be over - will call postRoundPageRender()
// Skip button - skip word (player2 had not guessed) and display new word
// next word - displays new word
func (page *gamePageObjects) roundPageRender() {

	page.prepareTimerForNewRound()

	page.prepareWordsForNewRound()

	layout := widgets.NewQGridLayout2()

	layout.AddWidget(page.forceQuitButton)
	layout.AddWidget(page.timerBoard)
	layout.AddWidget(page.wordLabel)
	layout.AddWidget(page.skipButton)
	layout.AddWidget(page.nextWordButton)

	page.application.show(layout)
}

// Set Team, who explains and who guessing names to labels
func (page *gamePageObjects) preparePreRoundLabels() {
	teamName, firstPlayerName, SecondPlayerName := page.game.GetCurrentPlayersName()

	page.teamNameLabel.SetText(teamName)
	page.firstPlayerLabel.SetText(firstPlayerName)
	page.secondPlayerLabel.SetText(SecondPlayerName)
}

// Render pre round page. Display Team, player1 and player2 names
func (page *gamePageObjects) preRoundPageRender() {
	page.preparePreRoundLabels()

	layout := widgets.NewQGridLayout2()

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

//================================================================================================================

// parent - window of app; teamName - raw array with names. NewGame([][]string) will work with it;
// language: 0 - Ru, 1 - En
func (app application) displayGameWindow(teamsName [][]string, language int) {
	// Connect close event, to prevent random quit
	app.window.ConnectCloseEvent(closeEvent)

	gamePage := newGamePage(app, teamsName, language)

	gamePage.render()
}
