package main

import (
	"Alias/internal/MainMenu"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Alias")
	window.SetMinimumSize2(850, 600)

	MainMenu.SetMainMenu(window)

	widgets.QApplication_Exec()
}
