package PreGameMenu

import (
	"github.com/therecipe/qt/widgets"
)

func changePushButtonToField(table *widgets.QTableWidget, row int, column int) {
	var widget = widgets.NewQTableWidgetItem(1)
	table.RemoveCellWidget(row, column)
	table.SetItem(row, column, widget)
}

func insertRow(table *widgets.QTableWidget, row int) {
	table.InsertRow(row)
	for i := 0; i < 3; i++ {

	}

	var addThirdPlayer = widgets.NewQPushButton(nil)
	addThirdPlayer.SetText("add 3 player")
	addThirdPlayer.ConnectPressed(func() {
		changePushButtonToField(table, row, 3)
	})
	table.SetCellWidget(row, 3, addThirdPlayer)
}

func PreGameMenu(parent *widgets.QMainWindow) {
	var centralWidget = widgets.NewQWidget(nil, 0)
	var layout = widgets.NewQGridLayout2()

	var title = widgets.NewQLabel(nil, 0)
	title.SetText("Pre Game")

	var backButton = widgets.NewQPushButton(nil)
	backButton.SetText("<")
	//TODO push on back button

	var teamTable = widgets.NewQTableWidget2(0, 4, nil)
	teamTable.SetMaximumWidth(1000)
	teamTable.SetMinimumWidth(830)
	namesForColumns := [4]string{"Team name", "1 player", "2 player", "3 player (optional)"}
	for i := 0; i < 4; i++ {
		teamTable.SetHorizontalHeaderItem(i, widgets.NewQTableWidgetItem2(namesForColumns[i], 1))
		teamTable.SetColumnWidth(i, 200)
	}
	teamTable.SetStyleSheet("QTableWidget::item {" +
		"border: 1px solid white;" +
		"}")

	insertRow(teamTable, 0)

	var teamsSpinBox = widgets.NewQSpinBox(nil)
	teamsSpinBox.SetMaximumWidth(50)
	teamsSpinBox.SetValueDefault(1)
	teamsSpinBox.SetMinimum(1)
	teamsSpinBox.SetMaximum(9)
	teamsSpinBox.ConnectValueChanged(func(val int) {
		if val >= teamTable.RowCount() {
			for i := teamTable.RowCount(); i < val; i++ {
				insertRow(teamTable, i)
			}
		} else {
			for i := teamTable.RowCount(); i > val; i-- {
				teamTable.RemoveRow(i - 1)
			}
		}
	})

	var startGame = widgets.NewQPushButton(nil)
	startGame.SetText("Start Game")
	startGame.ConnectPressed(func() {

	})

	layout.AddWidget(backButton)
	layout.AddWidget(title)
	layout.AddWidget(teamsSpinBox)
	layout.AddWidget(teamTable)
	layout.AddWidget(startGame)

	centralWidget.SetLayout(layout)
	parent.SetCentralWidget(centralWidget)
	parent.Show()
}
