package view

import (
	"github.com/andlabs/ui"
)

type Gui struct {
	initialized bool
	nbLig       int
	nbCol       int
	grid        [][]*ui.Checkbox
	gridData    [][]bool
}

func (v *Gui) Print(g [][]bool) {
	v.gridData = g
	if !v.initialized {
		v.initialized = true
		go v.firstPrint()
	}

	// refresh GUI
}

func (v *Gui) firstPrint() {
	v.nbLig = len(v.gridData)
	v.nbCol = len(v.gridData[0])

	v.grid = make([][]*ui.Checkbox, v.nbLig)
	for i := range v.grid {
		v.grid[i] = make([]*ui.Checkbox, v.nbCol)
	}

	err := ui.Main(func() {
		// create grid
		grid := ui.NewVerticalBox()
		for lig, ligSlice := range v.grid {
			lineUI := ui.NewHorizontalBox()
			for col := range ligSlice {
				btn := ui.NewCheckbox("")
				btn.OnToggled(func(*ui.Checkbox) {
					// Update data.Grid
					v.gridData[lig][col] = btn.Checked()
				})
				lineUI.Append(btn, false)
			}
			grid.Append(lineUI, false)
		}
		window := ui.NewWindow("Game of Life", v.nbCol, v.nbLig, false)
		window.SetChild(grid)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

func (v *Gui) AddCell() (lig, col int) {

	return
}
