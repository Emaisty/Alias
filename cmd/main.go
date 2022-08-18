package main

import (
	"Alias/internal/Render"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Alias_qt")
	window.SetMinimumSize2(850, 600)

	Render.RenderMainMenu(window)

	widgets.QApplication_Exec()
}
