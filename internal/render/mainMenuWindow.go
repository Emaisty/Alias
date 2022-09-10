package render

import (
	"Alias/internal/glob"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
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

	page.startGameButton.SetFixedSize(core.NewQSize2(300, 75))
}

func (page *mainMenuObjects) createWidgets() {
	page.title = widgets.NewQLabel2("Alias", nil, 0)

	// Create font style for label
	font := gui.NewQFont()
	font.SetPointSize(72)

	page.title.SetFont(font)

	page.createStartGameButton()
}

// Construct layout from widgets and render it on main window
func (page *mainMenuObjects) render() {

	//create widgets
	page.createWidgets()

	layout := widgets.NewQGridLayout2()

	layout.AddWidget2(page.title, 0, 0, core.Qt__AlignCenter|core.Qt__AlignTop)
	layout.AddWidget2(page.startGameButton, 1, 0, core.Qt__AlignCenter)

	page.application.show(layout)
}

//================================================================================================================

func (app application) displayMainMenuWindow() {
	//create main menu objects
	mainMenu := newMainMenuObjects(app)

	// Construct layout and render it on main window
	mainMenu.render()
}
