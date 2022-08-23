package render

import (
	"Alias/internal/glob"
	"encoding/json"
	"fmt"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"io/ioutil"
	"os"
)

//================================================================================================================
// Config block

// Struct to open and save config file
type configurations struct {
	Language       string
	Version        int
	CostOfGuessing int
	CostOfSkip     int
	TimeOfRound    int
	TargetScore    int
}

func readBytesFromFile(pathToFile string) []byte {
	//open file
	jsonFile, err := os.Open(pathToFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer jsonFile.Close()

	//read bytes from file
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return byteValue
}

func setGlobalConfiguration(config configurations) {
	glob.Config.CostOfGuessing = config.CostOfGuessing

	glob.Config.CostOfSkip = config.CostOfSkip

	glob.Config.TimeOfRound = config.TimeOfRound

	glob.Config.TargetScore = config.TargetScore
}

// Load configurations from file and set it to application. Read all phrases and load them to glob.Text
func initConfigs() {
	//load settings from json
	var config configurations
	err := json.Unmarshal(readBytesFromFile("data/config.json"), &config)
	if err != nil {
		return
	}

	//set configurations for application
	setGlobalConfiguration(config)

	//read all text from json
	var text glob.AllText
	err = json.Unmarshal(readBytesFromFile("data/text.json"), &text)
	if err != nil {
		return
	}

	//according to settings, set language
	switch config.Language {
	case "Ru":
		glob.Text = text.Ru
		break
	case "En":
		glob.Text = text.En
		break
	}

}

//================================================================================================================
// Class block

// Class of application
// Call NewApplication, when want to start program
// .Run() will start application and draw mainMenuWindow
type application struct {
	window *widgets.QMainWindow
}

// NewApplication create game Alias.
// To start in call Run()
func NewApplication() *application {
	// Load settings from file
	initConfigs()

	var app application
	widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Alias")
	window.SetMinimumSize2(850, 600)
	app.window = window

	return &app
}

//================================================================================================================
// methods for class

// Function to catch close event. Connect it to window.connectCloseEvent, to prevent user's random exit
func closeEvent(event *gui.QCloseEvent) {
	res := widgets.QMessageBox_Question(nil, glob.Text.Quit, glob.Text.AUSureExit, widgets.QMessageBox__Close|widgets.QMessageBox__No|widgets.QMessageBox__Yes, widgets.QMessageBox__Yes)
	if res != widgets.QMessageBox__Yes {
		event.Ignore()
	}
}

// Draw your layout on a main window
func (app application) show(layout *widgets.QGridLayout) {
	centralWidget := widgets.NewQWidget(nil, 0)
	centralWidget.SetLayout(layout)
	app.window.SetCentralWidget(centralWidget)
	app.window.Show()
}

//================================================================================================================

// Run will create window with a game
func (app application) Run() {
	app.displayMainMenuWindow()
	widgets.QApplication_Exec()
}
