package Render

import (
	"Alias/internal/glob"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func changePushButtonToField(table *widgets.QTableWidget, row int, column int) {
	widget := widgets.NewQTableWidgetItem(1)
	table.RemoveCellWidget(row, column)
	table.SetItem(row, column, widget)
}

func insertRow(table *widgets.QTableWidget, row int) {
	table.InsertRow(row)
	for i := 0; i < table.ColumnCount(); i++ {
		table.SetItem(row, i, widgets.NewQTableWidgetItem(1))
	}
	if table.ColumnCount() == 4 {
		addThirdPlayer := widgets.NewQPushButton(nil)
		addThirdPlayer.SetText(glob.Text.AddThirdPlayer)
		addThirdPlayer.ConnectPressed(func() {
			changePushButtonToField(table, row, 3)
		})
		table.SetCellWidget(row, 3, addThirdPlayer)
	}

}

// Check content of cell. If it is empty - colors it into red and return false
func checkCell(cell *widgets.QTableWidgetItem) bool {
	if cell.Text() == "" {
		cell.SetBackground(gui.NewQBrush4(core.Qt__red, core.Qt__SolidPattern))
		return false
	}
	return true
}

// Checks table for user-filled data and convert it into array of teams.
// return true, if everything OK and returns array of teams
// if returns false, something in table not full-filled.
func getTeamFromTable(table *widgets.QTableWidget) ([][]string, bool) {
	var (
		teams = make([][]string, 0)
		flag  = true
	)

	for i := 0; i < table.RowCount(); i++ {
		teams = append(teams, []string{})
		for j := 0; j < table.ColumnCount(); j++ {
			if j != 3 {
				flag = checkCell(table.Item(i, j)) && flag
			}
			teams[i] = append(teams[i], table.Item(i, j).Text())
		}
	}
	if flag {
		return teams, true
	}
	return nil, false
}

// Add/delete rows from table, when changed value on a spin box
func spinBoxValueChanged(val int, teamTable *widgets.QTableWidget) {
	if val >= teamTable.RowCount() {
		for i := teamTable.RowCount(); i < val; i++ {
			insertRow(teamTable, i)
		}
	} else {
		for i := teamTable.RowCount(); i > val; i-- {
			teamTable.RemoveRow(i - 1)
		}
	}
}

func createTable() *widgets.QTableWidget {
	teamTable := widgets.NewQTableWidget2(0, 4, nil)
	teamTable.SetMinimumHeight(1000)
	teamTable.SetMinimumWidth(830)
	teamTable.ConnectCellPressed(func(row int, column int) {
		teamTable.Item(row, column).SetBackground(gui.NewQBrush4(core.Qt__red, core.Qt__NoBrush))
	})
	namesForColumns := [4]string{glob.Text.TeamName, glob.Text.FirstPlayer, glob.Text.SecondPlayer,
		glob.Text.ThirdPlayer}
	for i := 0; i < 4; i++ {
		teamTable.SetHorizontalHeaderItem(i, widgets.NewQTableWidgetItem2(namesForColumns[i], 1))
		teamTable.SetColumnWidth(i, 200)
	}

	insertRow(teamTable, 0)
	insertRow(teamTable, 1)
	return teamTable
}

func createSpinBox(table *widgets.QTableWidget) *widgets.QSpinBox {
	teamsSpinBox := widgets.NewQSpinBox(nil)
	teamsSpinBox.SetMaximumWidth(50)
	teamsSpinBox.SetValueDefault(1)
	teamsSpinBox.SetMinimum(2)
	teamsSpinBox.SetMaximum(9)
	teamsSpinBox.ConnectValueChanged(func(val int) {
		spinBoxValueChanged(val, table)
	})
	return teamsSpinBox
}

func prepareTableForTeamMode(table *widgets.QTableWidget) {
	for i := 0; i < 3; i++ {
		table.InsertColumn(i + 1)
	}
	for i := 0; i < table.RowCount(); i++ {
		table.RemoveRow(i)
		insertRow(table, i)
	}
	namesForColumns := [4]string{glob.Text.TeamName, glob.Text.FirstPlayer, glob.Text.SecondPlayer,
		glob.Text.ThirdPlayer}
	for i := 0; i < 4; i++ {
		table.SetHorizontalHeaderItem(i, widgets.NewQTableWidgetItem2(namesForColumns[i], 1))
		table.SetColumnWidth(i, 200)
	}
}

func prepareTableForSoloMode(table *widgets.QTableWidget) {
	for i := 0; i < 3; i++ {
		table.RemoveColumn(1)
	}
	for i := 0; i < table.RowCount(); i++ {
		table.RemoveRow(i)
		insertRow(table, i)
	}
	table.SetHorizontalHeaderItem(0, widgets.NewQTableWidgetItem2(glob.Text.PlayerName, 1))
	table.SetColumnWidth(0, 200)
}

func contentOfPreGameMenu(parent *widgets.QMainWindow) *widgets.QGridLayout {
	layout := widgets.NewQGridLayout2()

	backButton := widgets.NewQPushButton(nil)
	backButton.SetText("<")
	backButton.ConnectPressed(func() {
		DisplayMainMenu(parent)
	})

	title := widgets.NewQLabel2(glob.Text.PreGame, nil, 0)

	labelHowManyTeams := widgets.NewQLabel2(glob.Text.HowManyTeams, nil, 0)

	teamTable := createTable()

	teamsSpinBox := createSpinBox(teamTable)

	modeComboBox := widgets.NewQComboBox(nil)
	modeComboBox.AddItems([]string{glob.Text.TeamMode, glob.Text.SoloMode})
	modeComboBox.ConnectCurrentIndexChanged(func(index int) {
		if index == 0 {
			labelHowManyTeams.SetText(glob.Text.HowManyTeams)
			prepareTableForTeamMode(teamTable)
		} else {
			labelHowManyTeams.SetText(glob.Text.HowManyPlayers)
			prepareTableForSoloMode(teamTable)
		}
	})

	languageComboBox := widgets.NewQComboBox(nil)
	languageComboBox.AddItems([]string{glob.Text.Russian, glob.Text.English})

	difficultyLabel := widgets.NewQLabel2(glob.Text.Difficulty, nil, 0)

	difficultyComboBox := widgets.NewQComboBox(nil)
	difficultyComboBox.AddItems([]string{"1", "2", "3"})

	startGame := widgets.NewQPushButton(nil)
	startGame.SetText(glob.Text.StartGame)
	startGame.ConnectPressed(func() {
		teams, flag := getTeamFromTable(teamTable)
		if !flag {
			return
		}
		_ = teams
		//TODO next
	})

	layout.AddWidget(backButton)
	layout.AddWidget(title)
	layout.AddWidget(modeComboBox)
	layout.AddWidget(labelHowManyTeams)
	layout.AddWidget(teamsSpinBox)
	layout.AddWidget(languageComboBox)
	layout.AddWidget(difficultyLabel)
	layout.AddWidget(difficultyComboBox)
	layout.AddWidget(teamTable)
	layout.AddWidget(startGame)

	return layout
}

func displayPreGameMenu(parent *widgets.QMainWindow) {
	centralWidget := widgets.NewQWidget(nil, 0)

	centralWidget.SetLayout(contentOfPreGameMenu(parent))
	parent.SetCentralWidget(centralWidget)
	parent.Show()
}
