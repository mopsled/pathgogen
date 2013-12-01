package main

import (
	"testing"
)

func TestCreateEmptyGrid(t *testing.T) {
	g := NewEmptyGrid(10, 10)
	if g == nil {
		t.Error("Can't create an empty grid")
	}
}

func TestEmptyGridHasEmptyCells(t *testing.T) {
	g := NewEmptyGrid(10, 10)
	for _, cells := range g.cells {
		for _, cell := range cells {
			if cell.cellType != nilCell {
				t.Error("Empty grid contains a non-empty cell")
			}
		}
	}
}

func TestCreateGridWithCells(t *testing.T) {
	cells := [][]Cell{
		[]Cell{Cell{nilCell, ""}, Cell{nilCell, ""}, Cell{nilCell, ""}},
		[]Cell{Cell{nilCell, ""}, Cell{aCell, ""}, Cell{nilCell, ""}},
		[]Cell{Cell{nilCell, ""}, Cell{nilCell, ""}, Cell{nilCell, ""}},
	}
	g := NewGridFromCells(cells)
	if g.cells[0][0].cellType != nilCell {
		t.Error("Unexpected cellType in grid initialized with cells")
	}
	if g.cells[1][1].cellType != aCell {
		t.Error("Unexpected cellType in grid initialized with cells")
	}
}

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

func TestAccessGridByCoordinateWithASmallBoard(t *testing.T) {
	cells := []string{
		"a...a",
		".b.b.",
		"..c..",
	}
	g, err := NewGridFromAscii(cells)

	cell, err := g.At("a3")
	if err != nil {
		t.Errorf("Grid access by coordinate returned an error: %s", err)
	}
	if cell.cellType != aCell {
		t.Error("Grid access by coordinate returned unexpected cellType")
	}

	cell, err = g.At("c1")
	if err != nil {
		t.Errorf("Grid access by coordinate returned an error: %s", err)
	}
	if cell.cellType != cCell {
		t.Error("Grid access by coordinate returned unexpected cellType")
	}
}

func TestAccessGridByCoordinateWithAGiantBoard(t *testing.T) {
	cells := []string{
		".................................................w",
		"..................................................",
		".......c..........................................",
		"..................................................",
		"..................................................",
		"..................................................",
		"..................................................",
		"..................................................",
		"...............................b..................",
		"..................................................",
		"..................................................",
		"..................................................",
		"..................................................",
		"..................................................",
		"..................................................",
		"..................................................",
		"a.................................................",
	}
	g, err := NewGridFromAscii(cells)
	cell, err := g.At("a1")
	if err != nil {
		t.Errorf("Grid access by coordinate returned an error: %s", err)
	}
	if cell.cellType != aCell {
		t.Error("Grid access by coordinate returned unexpected cellType")
	}

	cell, err = g.At("ax17")
	if err != nil {
		t.Errorf("Grid access by coordinate returned an error: %s", err)
	}
	if cell.cellType != wallCell {
		t.Error("Grid access by coordinate returned unexpected cellType")
	}

	if cell, _ := g.At("af9"); cell.cellType != bCell {
		t.Error("Grid access by coordinate returned unexpected cellType")
	}
	if cell, _ := g.At("h15"); cell.cellType != cCell {
		t.Error("Grid access by coordinate returned unexpected cellType")
	}
	if cell, _ := g.At("d12"); cell.cellType != nilCell {
		t.Error("Grid access by coordinate returned unexpected cellType")
	}
}

func TestAccessGridReturnsErrorForInvalidIndices(t *testing.T) {
	cells := []string{
		"a...a",
		".b.b.",
		"..c..",
	}
	g, _ := NewGridFromAscii(cells)

	if _, err := g.At("abc"); err == nil {
		t.Errorf("Grid access by coordinate should have an error for index 'abc', but didn't")
	}
	if _, err := g.At("1234"); err == nil {
		t.Errorf("Grid access by coordinate should have an error for index 'abc', but didn't")
	}
	if _, err := g.At("aa11"); err == nil {
		t.Errorf("Grid access by coordinate should have an error for index 'abc', but didn't")
	}
	if _, err := g.At("1a"); err == nil {
		t.Errorf("Grid access by coordinate should have an error for index 'abc', but didn't")
	}
	if _, err := g.At("a0"); err == nil {
		t.Errorf("Grid access by coordinate should have an error for index 'abc', but didn't")
	}
	if _, err := g.At("1z"); err == nil {
		t.Errorf("Grid access by coordinate should have an error for index 'abc', but didn't")
	}
}

func TestGridSet(t *testing.T) {
	cells := []string{
		"a...a..",
		".b.b...",
		"..c....",
	}
	g, _ := NewGridFromAscii(cells)

	err := g.Set("a1", teamlessCell(aCell))
	if err != nil {
		t.Errorf("Grid set by coordinate returned an error: %s", err)
	}
	if cell, _ := g.At("a1"); cell.cellType != aCell {
		t.Error("Grid set by coordinate returned unexpected cellType")
	}

	err = g.Set("a3", teamlessCell(bCell))
	if err != nil {
		t.Errorf("Grid set by coordinate returned an error: %s", err)
	}
	if cell, _ := g.At("a3"); cell.cellType != bCell {
		t.Error("Grid set by coordinate returned unexpected cellType")
	}

	err = g.Set("c2", teamlessCell(cCell))
	if err != nil {
		t.Errorf("Grid set by coordinate returned an error: %s", err)
	}
	if cell, _ := g.At("c2"); cell.cellType != cCell {
		t.Error("Grid set by coordinate returned unexpected cellType")
	}

	err = g.Set("g3", teamlessCell(nilCell))
	if err != nil {
		t.Errorf("Grid set by coordinate returned an error: %s", err)
	}
	if cell, _ := g.At("g3"); cell.cellType != nilCell {
		t.Error("Grid set by coordinate returned unexpected cellType")
	}
}

func TestGridSetWithInvalidInput(t *testing.T) {
	cells := []string{
		"a...a",
		".b.b.",
		"..c..",
	}
	g, _ := NewGridFromAscii(cells)

	if err := g.Set("abc", teamlessCell(aCell)); err == nil {
		t.Errorf("Grid access by coordinate should have an error for index 'abc', but didn't")
	}
	if err := g.Set("aa11", teamlessCell(bCell)); err == nil {
		t.Errorf("Grid access by coordinate should have an error for index 'abc', but didn't")
	}
	if err := g.Set("f4", teamlessCell(nilCell)); err == nil {
		t.Errorf("Grid access by coordinate should have an error for index 'abc', but didn't")
	}
	if err := g.Set("zzzz1000", teamlessCell(cCell)); err == nil {
		t.Errorf("Grid access by coordinate should have an error for index 'abc', but didn't")
	}
}

func teamlessCell(cellType CellType) *Cell {
	return &Cell{cellType, ""}
}
