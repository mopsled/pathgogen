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
		[]Cell{Cell{nilCell, 0}, Cell{nilCell, 0}, Cell{nilCell, 0}},
		[]Cell{Cell{nilCell, 0}, Cell{aCell, 0}, Cell{nilCell, 0}},
		[]Cell{Cell{nilCell, 0}, Cell{nilCell, 0}, Cell{nilCell, 0}},
	}
	g := NewGridFromCells(cells)
	if g.cells[0][0].cellType != nilCell {
		t.Error("Unexpected cellType in grid initialized with cells")
	}
	if g.cells[1][1].cellType != aCell {
		t.Error("Unexpected cellType in grid initialized with cells")
	}
}

func TestAccessGridByCoordinateWithASmallBoard(t *testing.T) {
	cells := []string{
		"a...a",
		".b.b.",
		"..c..",
	}
	g, err := NewGridFromAscii(cells)
	if err != nil {
		t.Fatalf("Grid created from ascii produced an error: %s", err)
	}
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
	if err != nil {
		t.Fatalf("Grid created from ascii produced an error: %s", err)
	}
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
	g, err := NewGridFromAscii(cells)
	if err != nil {
		t.Fatalf("Grid created from ascii produced an error: %s", err)
	}
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
	g, err := NewGridFromAscii(cells)
	if err != nil {
		t.Fatalf("Grid created from ascii produced an error: %s", err)
	}

	err = g.Set("a1", teamOneCell(aCell))
	if err != nil {
		t.Errorf("Grid set by coordinate returned an error: %s", err)
	}
	if cell, _ := g.At("a1"); cell.cellType != aCell {
		t.Error("Grid set by coordinate returned unexpected cellType")
	}

	err = g.Set("a3", teamOneCell(bCell))
	if err != nil {
		t.Errorf("Grid set by coordinate returned an error: %s", err)
	}
	if cell, _ := g.At("a3"); cell.cellType != bCell {
		t.Error("Grid set by coordinate returned unexpected cellType")
	}

	err = g.Set("c2", teamOneCell(cCell))
	if err != nil {
		t.Errorf("Grid set by coordinate returned an error: %s", err)
	}
	if cell, _ := g.At("c2"); cell.cellType != cCell {
		t.Error("Grid set by coordinate returned unexpected cellType")
	}

	err = g.Set("g3", teamOneCell(nilCell))
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
	g, err := NewGridFromAscii(cells)
	if err != nil {
		t.Fatalf("Grid created from ascii produced an error: %s", err)
	}

	if err = g.Set("abc", teamOneCell(aCell)); err == nil {
		t.Errorf("Grid access by coordinate should have an error for index 'abc', but didn't")
	}
	if err = g.Set("aa11", teamOneCell(bCell)); err == nil {
		t.Errorf("Grid access by coordinate should have an error for index 'abc', but didn't")
	}
	if err = g.Set("f4", teamOneCell(nilCell)); err == nil {
		t.Errorf("Grid access by coordinate should have an error for index 'abc', but didn't")
	}
	if err = g.Set("zzzz1000", teamOneCell(cCell)); err == nil {
		t.Errorf("Grid access by coordinate should have an error for index 'abc', but didn't")
	}
}

func teamOneCell(cellType CellType) *Cell {
	return &Cell{cellType, 1}
}
