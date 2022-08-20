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

// Struct to open and save config file
type configurations struct {
	Language    string
	Version     int
	CostOfSkip  int
	TimeOfRound int
}

func readBytesFromFile(pathToFile string) []byte {
	//open file
	jsonFile, err := os.Open(pathToFile)
	if err != nil {
		fmt.Println(err)
	}

	//read bytes from file
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	err = jsonFile.Close()
	if err != nil {
		return nil
	}

	return byteValue
}

func initConfigs() {
	//load settings from json
	var config configurations
	json.Unmarshal(readBytesFromFile("data/config.json"), &config)

	//read all text from json
	var text glob.AllText
	json.Unmarshal(readBytesFromFile("data/text.json"), &text)

	//according to settings, set language
	switch config.Language {
	case "Ru":
		glob.Text = text.Ru
		break
	case "En":
		glob.Text = text.En
		break
	}

	glob.Config.CostOfSkip = config.CostOfSkip

	glob.Config.TimeOfRound = config.TimeOfRound

}

type page interface {
	createObjects()
	render()
}

type application struct {
	window *widgets.QMainWindow
}

// Are you sure you want to exit?
func closeEvent(event *gui.QCloseEvent) {
	res := widgets.QMessageBox_Question(nil, glob.Text.Quit, glob.Text.AUSureExit, widgets.QMessageBox__Close|widgets.QMessageBox__No|widgets.QMessageBox__Yes, widgets.QMessageBox__Yes)
	if res != widgets.QMessageBox__Yes {
		event.Ignore()
	}
}

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

// Draw your layout on a main window
func (app application) show(layout *widgets.QGridLayout) {
	centralWidget := widgets.NewQWidget(nil, 0)
	centralWidget.SetLayout(layout)
	app.window.SetCentralWidget(centralWidget)
	app.window.Show()
}

// Run application
func (app application) Run() {
	app.displayMainMenu()
	widgets.QApplication_Exec()
}
