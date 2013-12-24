package main

import (
	"testing"
)

func TestCreateGridWithAscii(t *testing.T) {
	cells := []string{
		"a...a",
		".b.b.",
		"..c..",
	}
	g, err := NewGridFromAscii(cells)
	if err != nil {
		t.Errorf("Grid created from ascii produced an error: %s", err)
	}
	if g.Width() != 5 {
		t.Error("Grid has the wrong width for its initializing string")
	}
	if g.Height() != 3 {
		t.Error("Grid has the wrong width for its initializing string")
	}
	if g.cells[0][0].cellType != aCell {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if g.cells[2][2].cellType != cCell {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if g.cells[0][2].cellType != nilCell {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if g.cells[1][3].cellType != bCell {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
}

func TestStringForGridWithASimpleGrid(t *testing.T) {
	cells := []string{
		".....",
		".....",
		"..a..",
		".....",
		".....",
	}
	g, _ := NewGridFromAscii(cells)
	s, err := StringForGrid(g)
	if err != nil {
		t.Errorf("Error creating a string from grid: %s", err)
	}
	expectedString := ".....\n.....\n..a..\n.....\n.....\n"
	if s != expectedString {
		t.Errorf("Unexpected output for grid. Expected:\n%sGot:\n%s", expectedString, s)
	}
}

func TestStringForGridWithAModeratelyComplexGrid(t *testing.T) {
	cells := []string{
		"..a......",
		"b..ww....",
		"..aww.a..",
		"..bwwaba.",
		".bcb..b..",
		"..b......",
	}

	g, _ := NewGridFromAscii(cells)
	s, err := StringForGrid(g)
	if err != nil {
		t.Errorf("Error creating a string from grid: %s", err)
	}
	expectedString := "..a......\nb..ww....\n..aww.a..\n..bwwaba.\n.bcb..b..\n..b......\n"
	if s != expectedString {
		t.Errorf("Unexpected output for grid. Expected:\n%sGot:\n%s", expectedString, s)
	}
}
