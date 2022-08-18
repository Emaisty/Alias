package main

import (
	"Alias/internal/Render"
	"Alias/internal/glob"
	"encoding/json"
	"fmt"
	"github.com/therecipe/qt/widgets"
	"io/ioutil"
	"os"
)

type configurations struct {
	Language string
	Version  int
}

func initConfigs() {
	jsonFile, err := os.Open("data/config.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	jsonFile.Close()

	var config configurations
	json.Unmarshal([]byte(byteValue), &config)

	jsonFile, err = os.Open("data/text.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	jsonFile.Close()

	var text glob.AllText
	json.Unmarshal([]byte(byteValue), &text)

	switch config.Language {
	case "Ru":
		glob.Text = text.Ru
		break
	case "En":
		glob.Text = text.En
		break
	}

}

func main() {
	initConfigs()
	widgets.NewQApplication(len(os.Args), os.Args)

	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Alias_qt")
	window.SetMinimumSize2(850, 600)

	Render.DisplayMainMenu(window)

	widgets.QApplication_Exec()
}
