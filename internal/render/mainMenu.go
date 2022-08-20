package render

import (
	"Alias/internal/glob"
	"github.com/therecipe/qt/widgets"
)

type mainMenuObjects struct {
	application
	title           *widgets.QLabel
	startGameButton *widgets.QPushButton
}

func NewMainMenuObjects(app application) *mainMenuObjects {
	var mainMenu mainMenuObjects
	mainMenu.application = app
	return &mainMenu
}

func (page *mainMenuObjects) createTitle() {
	page.title = widgets.NewQLabel(nil, 0)
	page.title.SetText("Alias")
}

func (page *mainMenuObjects) createStartGameButton() {
	page.startGameButton = widgets.NewQPushButton(nil)
	page.startGameButton.SetText(glob.Text.StartGame)

	//call pre game page, when button pressed
	page.startGameButton.ConnectPressed(page.application.displayPreGameMenu)
}

func (page *mainMenuObjects) createObjects() {
	page.createTitle()
	page.createStartGameButton()
}

// Construct layout from elements and render it on main window
func (page *mainMenuObjects) render() {
	layout := widgets.NewQGridLayout2()
	layout.AddWidget(page.title)
	layout.AddWidget(page.startGameButton)

	page.application.show(layout)
}

func (app application) displayMainMenu() {
	//create main menu objects
	mainMenu := NewMainMenuObjects(app)

	//create widgets
	mainMenu.createObjects()

	// Construct layout and render it on main window
	mainMenu.render()
}
