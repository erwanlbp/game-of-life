package view

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

type Console struct {
}

func (c Console) Print(g [][]bool) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for _, line := range g {
		for _, cell := range line {
			if cell {
				fmt.Print("o ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func (c Console) AddCell() (lig, col int) {
	lig = rand.Intn(55)
	col = rand.Intn(100)
	return
}
