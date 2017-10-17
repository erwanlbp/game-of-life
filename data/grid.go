package data

import "github.com/erwanlbp/game-of-life/view"

type Grid [][]bool

func InitGrid(nbLig, nbCol int) Grid {
	grid := make(Grid, nbLig)
	for i := range grid {
		grid[i] = make([]bool, nbCol)
	}
	return grid
}

func (g Grid) Print() {
	view.View.Print([][]bool(g))
}

func (g Grid) AddCell() {
	lig, col := view.View.AddCell()
	if col < 0 || col > len(g[0])-1 || lig < 0 || lig > len(g)-1 {
		return
	}
	g[lig][col] = true
}

func (g Grid) Iterate() {
	newGrid := InitGrid(len(g), len(g[0]))

	for lig, line := range g {
		for col := range line {
			newGrid[lig][col] = g.NextIterationFor(lig, col)
		}
	}

	for lig, line := range g {
		for col := range line {
			g[lig][col] = newGrid[lig][col]
		}
	}
}

func (g Grid) NextIterationFor(lig, col int) bool {
	nbLivingAround := g.getNbLivingNeighbours(lig, col)
	if g[lig][col] {
		if nbLivingAround < 2 || nbLivingAround > 3 {
			return false
		}
	} else {
		if nbLivingAround == 3 {
			return true
		}
	}

	return g[lig][col]
}

func (g Grid) getNbLivingNeighbours(lig, col int) (nbAlive int) {
	for i := lig - 1; i <= lig+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i >= 0 && j >= 0 && i < len(g) && j < len(g[0]) && !(i == lig && j == col) && g[i][j] {
				nbAlive++
			}
		}
	}
	return
}
