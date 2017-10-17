package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/erwanlbp/game-of-life/data"
)

var nbLig, nbCol int

func init() {
	flag.IntVar(&nbLig, "nbLig", 55, "Number of lines")
	flag.IntVar(&nbCol, "nbCol", 100, "Number of columns")
	flag.Parse()
}

func main() {
	rand.Seed(time.Now().Unix())

	grid := data.InitGrid(nbLig, nbCol)

	grid.Print()

	for i := 0; i < 2000; i++ {
		grid.AddCell()
	}

	grid.Print()
	fmt.Println("Start in 3 sec ...")
	time.Sleep(time.Second * 3)

	start := time.Now()

	for i := 0; ; i++ {
		grid.Iterate()
		grid.Print()
		fmt.Print("Gen.", i)
		time.Sleep(time.Millisecond * 200)
	}

	elapsed := time.Since(start)
	fmt.Println("END : ", elapsed)
}
