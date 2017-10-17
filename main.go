package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/erwanlbp/game-of-life/data"
	"github.com/erwanlbp/game-of-life/view"
)

var nbLig, nbCol int

func init() {
	flag.IntVar(&nbLig, "nbLig", 30, "Number of lines")
	flag.IntVar(&nbCol, "nbCol", 30, "Number of columns")
	flag.Parse()
}

func main() {
	//rand.Seed(time.Now().Unix())

	grid := data.InitGrid(nbLig, nbCol)

	for i := 0; ; i++ {
		grid.Iterate()
		view.Console{}.Print([][]bool(grid))
		grid.Print()
		fmt.Println("Gen.", i)
		time.Sleep(time.Millisecond * 2000)
	}
}
