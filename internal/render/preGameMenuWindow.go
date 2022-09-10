package render

import (
	"Alias/internal/glob"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

//================================================================================================================
// Class block

type preGameMenuObjects struct {
	application
	backButton           *widgets.QPushButton
	title                *widgets.QLabel
	modeComboBox         *widgets.QComboBox
	modeOfAGameLabel     *widgets.QLabel
	howManyTeamLabel     *widgets.QLabel
	languageOfAGameLabel *widgets.QLabel
	difficultyLabel      *widgets.QLabel
	howManyTeamSpinBox   *widgets.QSpinBox
	languageComboBox     *widgets.QComboBox
	difficultyComboBox   *widgets.QComboBox
	startGameButton      *widgets.QPushButton
	table                *widgets.QTableWidget
}

func newPreGameMenuObjects(app application) *preGameMenuObjects {
	var menu preGameMenuObjects
	menu.application = app
	return &menu
}

//================================================================================================================

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
			if table.Item(i, j).Text() != "" {
				teams[i] = append(teams[i], table.Item(i, j).Text())
			}
		}
	}
	if flag {
		return teams, true
	}
	return nil, false
}

// Change rows according to a game mode
func (page *preGameMenuObjects) changeRows() {
	for i := 0; i < page.table.RowCount(); i++ {
		page.table.RemoveRow(i)
		insertRow(page.table, i)
	}
}

// Change table for team mode. Add 3 columns (4 in total)
func (page *preGameMenuObjects) prepareTableForTeamMode() {
	// Add columns for Player's names
	for i := 0; i < 3; i++ {
		page.table.InsertColumn(i + 1)
	}

	page.changeRows()

	// Set names for a columns
	namesForColumns := glob.Text.ColumnsForTeamMode
	for i := 0; i < 4; i++ {
		page.table.SetHorizontalHeaderItem(i, widgets.NewQTableWidgetItem2(namesForColumns[i], 1))
		//page.table.SetColumnWidth(i, 200)
	}
}

// Change table for solo mode. Left only 1 column
func (page *preGameMenuObjects) prepareTableForSoloMode() {
	// Remove 2-4 columns
	for i := 0; i < 3; i++ {
		page.table.RemoveColumn(1)
	}

	page.changeRows()

	// Set name for columns
	page.table.SetHorizontalHeaderItem(0, widgets.NewQTableWidgetItem2(glob.Text.PlayerName, 1))
	//page.table.SetColumnWidth(0, 200)
}

//================================================================================================================
// Events

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

// Event, when (row,column) cell in table pressed
func (page *preGameMenuObjects) cellPressed(row int, column int) {
	page.table.Item(row, column).SetBackground(gui.NewQBrush4(core.Qt__red, core.Qt__NoBrush))
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

//================================================================================================================
// Create and render block

func (page *preGameMenuObjects) createComboBoxes() {
	page.modeComboBox = widgets.NewQComboBox(nil)
	page.modeComboBox.AddItems([]string{glob.Text.TeamMode, glob.Text.SoloMode})
	page.modeComboBox.ConnectCurrentIndexChanged(page.modeChanged)

	page.languageComboBox = widgets.NewQComboBox(nil)
	page.languageComboBox.AddItems([]string{glob.Text.Russian, glob.Text.English})

	page.difficultyComboBox = widgets.NewQComboBox(nil)
	page.difficultyComboBox.AddItems([]string{"1", "2", "3"})
}

// Spin Box connected to table and when value changed - changing count of rows
func (page *preGameMenuObjects) createSpinBox() {
	page.howManyTeamSpinBox = widgets.NewQSpinBox(nil)
	//page.howManyTeamSpinBox.SetMaximumWidth(50)
	page.howManyTeamSpinBox.SetFixedWidth(50)
	page.howManyTeamSpinBox.SetValueDefault(1)
	page.howManyTeamSpinBox.SetMinimum(2)
	page.howManyTeamSpinBox.SetMaximum(9)

	// Connect change event
	page.howManyTeamSpinBox.ConnectValueChanged(page.spinBoxValueChanged)
}

func (page *preGameMenuObjects) createTable() {

	// Create table and set parameters to it
	page.table = widgets.NewQTableWidget2(0, 4, nil)
	//page.table.SetMinimumHeight(1000)
	//page.table.SetMinimumWidth(830)
	page.table.ConnectCellPressed(page.cellPressed)

	// Set names to table columns
	namesForColumns := glob.Text.ColumnsForTeamMode
	for i := 0; i < 4; i++ {
		page.table.SetHorizontalHeaderItem(i, widgets.NewQTableWidgetItem2(namesForColumns[i], 1))
		//page.table.SetColumnWidth(i, 200)
	}

	page.table.HorizontalHeader().SetSectionResizeMode(widgets.QHeaderView__Stretch)
	// Insert first 2 rows into table
	insertRow(page.table, 0)
	insertRow(page.table, 1)
}

// Create Back button (return to main page) and title for page
func (page *preGameMenuObjects) createHeader() {
	page.backButton = widgets.NewQPushButton(nil)
	page.backButton.SetText("<")
	page.backButton.SetFixedWidth(50)

	page.backButton.ConnectPressed(page.application.displayMainMenuWindow)

	page.title = widgets.NewQLabel2(glob.Text.PreGame, nil, 0)
	// Create font style for label
	font := gui.NewQFont()
	font.SetPointSize(42)

	page.title.SetFont(font)

	// Create labels. Uses only for more friendly UI
	page.modeOfAGameLabel = widgets.NewQLabel2(glob.Text.GameMode, nil, 0)

	page.howManyTeamLabel = widgets.NewQLabel2(glob.Text.HowManyTeams, nil, 0)

	page.languageOfAGameLabel = widgets.NewQLabel2(glob.Text.Language, nil, 0)

	page.difficultyLabel = widgets.NewQLabel2(glob.Text.Difficulty, nil, 0)
}

// Create table for teams/players names, buttons to control modes etc.
func (page *preGameMenuObjects) createBodyOfPage() {

	// Create table with teams/players names
	page.createTable()

	// Spin box. Connected with number of table rows. Increase/decrease their count
	page.createSpinBox()

	// Combo box with mode choose, language and difficulty
	page.createComboBoxes()

	// Create Button to start a game
	page.startGameButton = widgets.NewQPushButton(nil)
	page.startGameButton.SetText(glob.Text.StartGame)
	page.startGameButton.ConnectPressed(page.runGame)
	page.startGameButton.SetFixedHeight(75)
}

func (page *preGameMenuObjects) createWidgets() {
	page.createHeader()

	page.createBodyOfPage()
}

func (page *preGameMenuObjects) render() {

	layout := widgets.NewQGridLayout2()

	// back button and title
	layout.AddWidget2(page.backButton, 0, 0, core.Qt__AlignLeft)
	layout.AddWidget2(page.title, 0, 2, 0)

	// Labels
	layout.AddWidget2(page.modeOfAGameLabel, 1, 0, core.Qt__AlignCenter)
	layout.AddWidget2(page.howManyTeamLabel, 1, 1, core.Qt__AlignCenter)
	layout.AddWidget2(page.languageOfAGameLabel, 1, 3, core.Qt__AlignCenter)
	layout.AddWidget2(page.difficultyLabel, 1, 4, core.Qt__AlignCenter)

	// Widgets with settings for a game
	layout.AddWidget2(page.modeComboBox, 2, 0, core.Qt__AlignCenter)
	layout.AddWidget2(page.howManyTeamSpinBox, 2, 1, core.Qt__AlignCenter)
	layout.AddWidget2(page.languageComboBox, 2, 3, core.Qt__AlignCenter)
	layout.AddWidget2(page.difficultyComboBox, 2, 4, core.Qt__AlignCenter)

	//table and button to start a game
	layout.AddWidget3(page.table, 3, 0, 1, 5, 0)
	layout.AddWidget2(page.startGameButton, 4, 2, 0)

	page.application.show(layout)
}

//================================================================================================================

func (app application) displayPreGameMenuWindow() {
	preGameMenu := newPreGameMenuObjects(app)

	preGameMenu.createWidgets()

	// Display all widgets
	preGameMenu.render()
}
