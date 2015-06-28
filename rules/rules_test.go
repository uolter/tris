package rules

import "testing"

func TestAddOK(t *testing.T) {

	tbl := new(Table)

	got := tbl.Add(0,0, 1)
	if got != nil {
		t.Error("Expected nil got ", got)
	}

	tbl.Show()

}

func TestAddKO(t *testing.T) {

	tbl := new(Table)

	tbl.Add(2,2,1)
	got := tbl.Add(2,2,1)
	if got == nil {
		t.Error("Cell Expected already taken. Got ", got)
	}

}

func TestAddKOVal(t *testing.T) {

	tbl := new(Table)

	got := tbl.Add(2,2,5)
	if got == nil {
		t.Error("Expected Error. Value not allowed.")
	}
}

func TestGetOk(t *testing.T) {

	tbl := new(Table)

	var got int

	got = tbl.Get(2,2)
	if got != 0 {
		t.Error("Cell Expected false got", got)
	}

	tbl.Add(2,2,1)

	got = tbl.Get(2,2)
	if got != 1 {
		t.Error("Cell Expected false got", got)
	}
}

func TestOneRowWin(t *testing.T) {

	tbl := new(Table)
	tbl.Add(0,0,1)
	tbl.Add(0,1,1)
	tbl.Add(0,2,1)

	win, _ := tbl.Win()

	if win  == false {
		t.Error("Expected one first row win.")
	}
}

func TestOneColsWin(t *testing.T) {

	tbl := new(Table)
	tbl.Add(0,1,2)
	tbl.Add(1,1,2)
	tbl.Add(2,1,2)

	win, _ := tbl.Win()

	if win == false {
		t.Error("Expected one first row win.")
	}

}


func TestCrossWin(t *testing.T) {

	tbl := new(Table)
	tbl.Add(0,0,2)
	tbl.Add(1,1,2)
	tbl.Add(2,2,2)

	win, _ := tbl.Win()

	if win == false {
		t.Error("Expected diagonal wins")
	}

	tbl = new(Table)

	win, _ = tbl.Win()

	if win == true {
		t.Error("Empty table. No winner expected.")
	}

	tbl.Add(0,2,1)
	tbl.Add(1,1,1)
	tbl.Add(2,0,1)

	win, _ = tbl.Win()

	if win == false {
		t.Error("Expected diagonal wins")
	}
}

func TestFreeCells(t *testing.T) {
	tbl := new(Table)
	freeCells := tbl.GetFreeCells()

	if len(freeCells) != 9 {
		t.Error("All cells expected free.")
	}

	tbl.Add(1,1,1)
	tbl.Add(2,1,2)
	freeCells = tbl.GetFreeCells()

	if len(freeCells) != 7 {
		t.Error("2 cells expected in use.")
	}

}