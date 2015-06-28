package main

import (
	"github.com/uolter/tris/rules"
	"fmt"
	"math/rand"
	)

var p = fmt.Print
var pln = fmt.Println

func main() {

	pln("Start playing")
	pln("O: You")
	pln("X: Go")
	t := new(rules.Table)

	var freeCells []rules.Cell 
	var cell rules.Cell

	for i:=0; i<9; i++ {
		
		freeCells = t.GetFreeCells()
		cell = freeCells[rand.Intn(len(freeCells))]

		t.Add(cell.Row, cell.Col, 1)
		t.Show()
		pln("")
		win, _ := t.Win()

		if win == true {
			pln("Go is the winner :-(")
			break
		}
		for {
			p("Row [0,1,2] ")
			fmt.Scanln(&cell.Row)
			p("Col [0,1,2] ")
			fmt.Scanln(&cell.Col)
			if t.Add(cell.Row, cell.Col, 2) == nil {
				break
			}
		}
		win, _ = t.Win()

		if win == true {
			pln("You win!!")
			break
		}
	}
}