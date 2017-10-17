package view

type view interface {
	Print(g [][]bool)
	AddCell() (int, int)
}

var View view = &Gui{}
