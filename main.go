package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andlabs/ui"
	"github.com/erwanlbp/game-of-life/data"
)

var nbLig, nbCol int

func init() {
	flag.IntVar(&nbLig, "nbLig", 55, "Number of lines")
	flag.IntVar(&nbCol, "nbCol", 100, "Number of columns")
	flag.Parse()
}

func main() {
	guiTEST()
	os.Exit(0)

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

func guiTEST() {
	err := ui.Main(func() {
		name := ui.NewEntry()
		button := ui.NewButton("Greet")
		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter your name:"), false)
		box.Append(name, false)
		box.Append(button, false)
		box.Append(greeting, false)
		window := ui.NewWindow("Hello", 200, 100, false)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			greeting.SetText("Hello, " + name.Text() + "!")
		})
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
