package render

import "github.com/therecipe/qt/widgets"

// parent - window of app; teamName - raw array with names. NewGame([][]string) will work with it;
// language: 0 - Ru, 1 - En
func (app application) displayGameWindow(teamsName [][]string, language int) {

	//create layout from widgets
	layout := widgets.NewQGridLayout2()

	//set layout to window and show
	app.show(layout)
}
