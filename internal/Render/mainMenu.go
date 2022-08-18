package Render

import (
	"Alias/internal/glob"
	"github.com/therecipe/qt/widgets"
)

func DisplayMainMenu(parent *widgets.QMainWindow) {
	centralWidget := widgets.NewQWidget(nil, 0)

	title := widgets.NewQLabel(nil, 0)
	title.SetText("Alias")

	var startGameButton = widgets.NewQPushButton(nil)
	startGameButton.SetText(glob.Text.StartGame)
	startGameButton.ConnectPressed(func() {
		displayPreGameMenu(parent)
	})

	layout := widgets.NewQGridLayout2()
	layout.AddWidget(title)
	layout.AddWidget(startGameButton)

	centralWidget.SetLayout(layout)
	parent.SetCentralWidget(centralWidget)
	parent.Show()
}
