package render

import (
	"Alias/internal/glob"
	"encoding/json"
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"io"
	"io/ioutil"
	"log"
	"os"
)

//================================================================================================================
// Config block

// Struct to open and save config file
type configurations struct {
	Language       string // language of a game
	Version        int    // version of a game
	CostOfGuessing int    // how much points you get, when guessing word
	CostOfSkip     int    // how much points you lose, when skip word
	TimeOfRound    int    // how much seconds round goes
	TargetScore    int    // after which amount of points game will end
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

// checkForConfigFile checking, if config file exist in /Library/Application Support. If not - copy it from inner files
func checkForConfigFile() {

	// if file exists - do nothing
	if _, err := os.Stat(glob.PathToUser + "/Library/Application Support/Alias/config.json"); err == nil {
		return
	}

	// if file do not exist - copy it from config.json in app dir
	srcFile, _ := os.Open(core.QCoreApplication_ApplicationDirPath() + "/../Resources/config.json")
	defer srcFile.Close()

	// Create dir for application config files
	os.MkdirAll(glob.PathToUser+"/Library/Application Support/Alias", os.ModePerm)

	// creates if file doesn't exist
	destFile, _ := os.Create(glob.PathToUser + "/Library/Application Support/Alias/config.json")
	defer destFile.Close()

	io.Copy(destFile, srcFile)
}

func setGlobalConfiguration(config configurations) {
	glob.Config.CostOfGuessing = config.CostOfGuessing

	glob.Config.CostOfSkip = config.CostOfSkip

	glob.Config.TimeOfRound = config.TimeOfRound

	glob.Config.TargetScore = config.TargetScore
}

// Load configurations from file and set it to application. Read all phrases and load them to glob.Text
func initConfigs() {

	// set path to home dir
	glob.PathToUser, _ = os.UserHomeDir()

	checkForConfigFile()

	//load settings from json
	var config configurations
	err := json.Unmarshal(readBytesFromFile(glob.PathToUser+"/Library/Application Support/Alias/config.json"), &config)
	if err != nil {
		log.Println(err)
		return
	}

	//set configurations for application
	setGlobalConfiguration(config)

	//read all text from json
	var text glob.AllText
	err = json.Unmarshal(readBytesFromFile(core.QCoreApplication_ApplicationDirPath()+"/../Resources/text.json"), &text)
	if err != nil {
		log.Println(err)
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
	var app application
	widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Alias")
	window.SetMinimumSize2(850, 600)
	app.window = window

	// Load settings from file
	initConfigs()

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
