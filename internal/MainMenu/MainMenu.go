package MainMenu

import (
	"Alias/internal/PreGameMenu"
	"github.com/therecipe/qt/widgets"
)

func SetMainMenu(parent *widgets.QMainWindow) {
	var centralWidget = widgets.NewQWidget(nil, 0)

	var title = widgets.NewQLabel(nil, 0)
	title.SetText("Alias")

	var startGameButton = widgets.NewQPushButton(nil)
	startGameButton.SetText("Start Game")
	startGameButton.ConnectPressed(func() {
		PreGameMenu.PreGameMenu(parent)
	})

	var layout = widgets.NewQGridLayout2()
	layout.AddWidget(title)
	layout.AddWidget(startGameButton)

	centralWidget.SetLayout(layout)
	parent.SetCentralWidget(centralWidget)
	parent.Show()
}
