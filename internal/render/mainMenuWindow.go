package render

import (
	"Alias/internal/glob"
	"github.com/therecipe/qt/widgets"
)

//================================================================================================================
// Class block

// widgets on main menu window
type mainMenuObjects struct {
	application
	title           *widgets.QLabel
	startGameButton *widgets.QPushButton
}

func newMainMenuObjects(app application) *mainMenuObjects {
	var mainMenu mainMenuObjects
	mainMenu.application = app
	return &mainMenu
}

//================================================================================================================
// Create and render block

func (page *mainMenuObjects) createStartGameButton() {
	page.startGameButton = widgets.NewQPushButton2(glob.Text.StartGame, nil)

	//call pre game window, when button pressed
	page.startGameButton.ConnectPressed(page.application.displayPreGameMenuWindow)
}

func (page *mainMenuObjects) createWidgets() {
	page.title = widgets.NewQLabel2("Alias", nil, 0)

	page.createStartGameButton()
}

// Construct layout from widgets and render it on main window
func (page *mainMenuObjects) render() {

	//create widgets
	page.createWidgets()

	layout := widgets.NewQGridLayout2()

	layout.AddWidget(page.title)
	layout.AddWidget(page.startGameButton)

	page.application.show(layout)
}

//================================================================================================================

func (app application) displayMainMenuWindow() {
	//create main menu objects
	mainMenu := newMainMenuObjects(app)

	// Construct layout and render it on main window
	mainMenu.render()
}
