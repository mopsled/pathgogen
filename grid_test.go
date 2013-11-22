package main

import (
	"testing"
)

func TestCreateEmptyBoard(t *testing.T) {
	g := NewEmptyGrid(10, 10)
	if g == nil {
		t.Error("Can't create an empty grid")
	}
}

func TestEmptyBoardHasEmptyCells(t *testing.T) {
	g := NewEmptyGrid(10, 10)
	for _, cells := range g.cells {
		for _, cell := range cells {
			if cell.cellType != nilCell {
				t.Error("Empty grid contains a non-empty cell")
			}
		}
	}
}

func TestCreateBoardWithCells(t *testing.T) {
	cells := [][]Cell{
		[]Cell{Cell{nilCell, ""}, Cell{nilCell, ""}, Cell{nilCell, ""}},
		[]Cell{Cell{nilCell, ""}, Cell{aCell, ""}, Cell{nilCell, ""}},
		[]Cell{Cell{nilCell, ""}, Cell{nilCell, ""}, Cell{nilCell, ""}},
	}
	g := NewGridFromCells(cells)
	if g.cells[0][0].cellType != nilCell {
		t.Error("Grid is different than the literal it was initialized with")
	}
	if g.cells[1][1].cellType != aCell {
		t.Error("Grid is different than the literal it was initialized with")
	}
}
