package rules

import (
	"errors"
	"fmt"
	)

var p = fmt.Print 

type Cell struct {
	Row, Col int
}

type Table struct {
	cells [3][3]int
}


func (t *Table) Add(row, col, val int) error {

	if t.cells[row][col] > 0 {
		return errors.New("Cell already set.") 
	}

	if val != 1 && val != 2 {
		return errors.New("Value must be 1 or 2.")
	}
			
	t.cells[row][col] = val
		
	return nil
}
		

func (t *Table) Get(row, col int) int {
	return t.cells[row][col]
}

func (t *Table) Show() {

	fmt.Println(" 012")
	for r:=0; r<3; r++ {
		p(r)
		for c:=0; c<3; c++ {
			if t.Get(r,c) == 0 {
				p("=")
			} else if t.Get(r,c) == 1 {
			   p("X")	
			} else if t.Get(r,c) == 2 {
				p("O")
			}
		}
		fmt.Println("")
	}
}

func (t *Table) Win() (bool, int) {

	for i:=0; i<3; i++ {
		// check all rows
		if (t.cells[i][0]!=0) && (t.cells[i][0] == t.cells[i][1]) && (t.cells[i][1] == t.cells[i][2]) {
			return true, t.cells[i][0]	
		}
		// check all columns.
		if (t.cells[0][i]!=0) && (t.cells[0][i] == t.cells[1][i]) && (t.cells[1][i] == t.cells[2][i]) {
			return true, t.cells[0][i]
		}
	} 

	// Test crossing:
	if (t.cells[0][0]!=0) &&  (t.cells[0][0] == t.cells[1][1]) && (t.cells[1][1] == t.cells[2][2]) {
			return true, t.cells[0][0]	
	}

	// Test crossing:
	if (t.cells[0][2]!=0) && (t.cells[0][2] == t.cells[1][1]) && (t.cells[1][1] == t.cells[2][0]) {
			return true, t.cells[0][2]	
	}

	return false, 0

}


func (t *Table) GetFreeCells() []Cell {
	var freeCells []Cell
	var cell *Cell
	for r:=0; r<3; r++  {
		for c:=0; c<3; c++ {
			if t.Get(r,c) == 0 {
				cell = new(Cell)
				cell.Row = r
				cell.Col = c
				freeCells = append(freeCells, *cell)
			}
		}
	}
	return freeCells
}