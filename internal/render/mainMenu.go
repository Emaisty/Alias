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

func (menu *mainMenuObjects) createTitle() {
	menu.title = widgets.NewQLabel(nil, 0)
	menu.title.SetText("Alias")
}

func (menu *mainMenuObjects) createStartGameButton() {
	menu.startGameButton = widgets.NewQPushButton(nil)
	menu.startGameButton.SetText(glob.Text.StartGame)

	//call pre game menu, when button pressed
	menu.startGameButton.ConnectPressed(menu.application.displayPreGameMenu)
}

// Construct layout from elements and render it on main window
func (menu *mainMenuObjects) render() {
	layout := widgets.NewQGridLayout2()
	layout.AddWidget(menu.title)
	layout.AddWidget(menu.startGameButton)

	menu.application.show(layout)
}

func (app application) displayMainMenu() {
	//create main menu objects
	mainMenu := NewMainMenuObjects(app)

	//set title
	mainMenu.createTitle()

	//button to run a game
	mainMenu.createStartGameButton()

	// Construct layout and render it on main window
	mainMenu.render()
}
