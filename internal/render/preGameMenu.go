package render

import (
	"Alias/internal/glob"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type preGameMenuObjects struct {
	application
	backButton         *widgets.QPushButton
	title              *widgets.QLabel
	modeComboBox       *widgets.QComboBox
	howManyTeamLabel   *widgets.QLabel
	howManyTeamSpinBox *widgets.QSpinBox
	languageComboBox   *widgets.QComboBox
	difficultyLabel    *widgets.QLabel
	difficultyComboBox *widgets.QComboBox
	table              *widgets.QTableWidget
	startGameButton    *widgets.QPushButton
}

func NewPreGameMenuObjects(app application) *preGameMenuObjects {
	var menu preGameMenuObjects
	menu.application = app
	return &menu
}

// Add/delete rows from table, when changed value on a spin box
func (page *preGameMenuObjects) spinBoxValueChanged(val int) {
	if val >= page.table.RowCount() {
		for i := page.table.RowCount(); i < val; i++ {
			insertRow(page.table, i)
		}
	} else {
		for i := page.table.RowCount(); i > val; i-- {
			page.table.RemoveRow(i - 1)
		}
	}
}

// Spin Box connected to table and when value changed - changing count of rows
func (page *preGameMenuObjects) createSpinBox() {
	page.howManyTeamSpinBox = widgets.NewQSpinBox(nil)
	page.howManyTeamSpinBox.SetMaximumWidth(50)
	page.howManyTeamSpinBox.SetValueDefault(1)
	page.howManyTeamSpinBox.SetMinimum(2)
	page.howManyTeamSpinBox.SetMaximum(9)
	page.howManyTeamSpinBox.ConnectValueChanged(page.spinBoxValueChanged)
}

// Event, when (row,column) cell in table pressed
func (page *preGameMenuObjects) cellPressed(row int, column int) {
	page.table.Item(row, column).SetBackground(gui.NewQBrush4(core.Qt__red, core.Qt__NoBrush))
}

func (page *preGameMenuObjects) createTable() {

	// Create table and set parameters to it
	page.table = widgets.NewQTableWidget2(0, 4, nil)
	page.table.SetMinimumHeight(1000)
	page.table.SetMinimumWidth(830)
	page.table.ConnectCellPressed(page.cellPressed)

	// Set names to table columns
	namesForColumns := [4]string{glob.Text.TeamName, glob.Text.FirstPlayer, glob.Text.SecondPlayer,
		glob.Text.ThirdPlayer}
	for i := 0; i < 4; i++ {
		page.table.SetHorizontalHeaderItem(i, widgets.NewQTableWidgetItem2(namesForColumns[i], 1))
		page.table.SetColumnWidth(i, 200)
	}

	// Insert first 2 rows into table
	insertRow(page.table, 0)
	insertRow(page.table, 1)
}

// Change table for team mode. Add 3 columns (4 in total)
func (page *preGameMenuObjects) prepareTableForTeamMode() {
	// Add columns for Player's names
	for i := 0; i < 3; i++ {
		page.table.InsertColumn(i + 1)
	}

	// Change rows according to a game mode
	for i := 0; i < page.table.RowCount(); i++ {
		page.table.RemoveRow(i)
		insertRow(page.table, i)
	}

	// Set names for a columns
	namesForColumns := [4]string{glob.Text.TeamName, glob.Text.FirstPlayer, glob.Text.SecondPlayer,
		glob.Text.ThirdPlayer}
	for i := 0; i < 4; i++ {
		page.table.SetHorizontalHeaderItem(i, widgets.NewQTableWidgetItem2(namesForColumns[i], 1))
		page.table.SetColumnWidth(i, 200)
	}
}

// Change table for solo mode. Left only 1 column
func (page *preGameMenuObjects) prepareTableForSoloMode() {
	// Remove 2-4 columns
	for i := 0; i < 3; i++ {
		page.table.RemoveColumn(1)
	}

	// Change rows according to a game mode
	for i := 0; i < page.table.RowCount(); i++ {
		page.table.RemoveRow(i)
		insertRow(page.table, i)
	}

	// Set name for columns
	page.table.SetHorizontalHeaderItem(0, widgets.NewQTableWidgetItem2(glob.Text.PlayerName, 1))
	page.table.SetColumnWidth(0, 200)
}

// When changed index - set another mode for table. 0 - team mode, 1 - solo mode
func (page *preGameMenuObjects) modeChanged(index int) {
	if index == 0 {
		page.howManyTeamLabel.SetText(glob.Text.HowManyTeams)
		page.prepareTableForTeamMode()
	} else {
		page.howManyTeamLabel.SetText(glob.Text.HowManyPlayers)
		page.prepareTableForSoloMode()
	}
}

// Event, when StartGameButton pressed
func (page *preGameMenuObjects) runGame() {
	// Get names of players/teams into array
	teams, flag := getTeamFromTable(page.table)
	if !flag {
		return
	}
	// Start a game
	page.application.displayGameWindow(teams, page.languageComboBox.CurrentIndex())
}

// Create Back button (return to main page) and title for page
func (page *preGameMenuObjects) createHeader() {
	page.backButton = widgets.NewQPushButton(nil)
	page.backButton.SetText("<")
	page.backButton.ConnectPressed(page.application.displayMainMenu)

	page.title = widgets.NewQLabel2(glob.Text.PreGame, nil, 0)
}

func (page *preGameMenuObjects) createBodyOfPage() {
	// Create labels
	page.howManyTeamLabel = widgets.NewQLabel2(glob.Text.HowManyTeams, nil, 0)

	page.difficultyLabel = widgets.NewQLabel2(glob.Text.Difficulty, nil, 0)

	// Create main objects
	page.createTable()

	page.createSpinBox()

	// Crate Combo Boxes
	page.modeComboBox = widgets.NewQComboBox(nil)
	page.modeComboBox.AddItems([]string{glob.Text.TeamMode, glob.Text.SoloMode})
	page.modeComboBox.ConnectCurrentIndexChanged(page.modeChanged)

	page.languageComboBox = widgets.NewQComboBox(nil)
	page.languageComboBox.AddItems([]string{glob.Text.Russian, glob.Text.English})

	page.difficultyComboBox = widgets.NewQComboBox(nil)
	page.difficultyComboBox.AddItems([]string{"1", "2", "3"})

	// Create Button to start a game
	page.startGameButton = widgets.NewQPushButton(nil)
	page.startGameButton.SetText(glob.Text.StartGame)
	page.startGameButton.ConnectPressed(page.runGame)
}

func (page *preGameMenuObjects) createObjects() {
	page.createHeader()

	page.createBodyOfPage()
}

func (page *preGameMenuObjects) render() {
	layout := widgets.NewQGridLayout2()

	layout.AddWidget(page.backButton)
	layout.AddWidget(page.title)
	layout.AddWidget(page.modeComboBox)
	layout.AddWidget(page.howManyTeamLabel)
	layout.AddWidget(page.howManyTeamSpinBox)
	layout.AddWidget(page.languageComboBox)
	layout.AddWidget(page.difficultyLabel)
	layout.AddWidget(page.difficultyComboBox)
	layout.AddWidget(page.table)
	layout.AddWidget(page.startGameButton)

	page.application.show(layout)
}

// Then pressed button in the 4th column - change button to LineEdit field
func changePushButtonToField(table *widgets.QTableWidget, row int, column int) {
	widget := widgets.NewQTableWidgetItem(1)
	table.RemoveCellWidget(row, column)
	table.SetItem(row, column, widget)
}

// Insert row into table. If there are 4 columns - 4th column will be a button
func insertRow(table *widgets.QTableWidget, row int) {
	// Insert raw
	table.InsertRow(row)
	// Insert cells
	for i := 0; i < table.ColumnCount(); i++ {
		table.SetItem(row, i, widgets.NewQTableWidgetItem(1))
	}
	// If there are 4 columns - insert pushButton into 4th column
	if table.ColumnCount() == 4 {
		addThirdPlayer := widgets.NewQPushButton(nil)
		addThirdPlayer.SetText(glob.Text.AddThirdPlayer)
		// Set action on button
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
		// If any field will be empty - flag will be false
		flag = true
	)

	for i := 0; i < table.RowCount(); i++ {
		teams = append(teams, []string{})
		for j := 0; j < table.ColumnCount(); j++ {
			// Check cell. If it is empty - colorized it. Except 4th - not mandatory column
			if j != 3 {
				flag = checkCell(table.Item(i, j)) && flag
			}
			// Push string from cell into array
			teams[i] = append(teams[i], table.Item(i, j).Text())
		}
	}
	if flag {
		return teams, true
	}
	return nil, false
}

func (app application) displayPreGameMenu() {
	preGameMenu := NewPreGameMenuObjects(app)

	// Create widgets for pre game page
	preGameMenu.createObjects()

	// Display all widgets
	preGameMenu.render()
}
