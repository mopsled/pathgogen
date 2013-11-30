package main

import (
	"errors"
)

type CellType int

const (
	nilCell CellType = iota
	aCell
	bCell
	cCell
	wallCell
)

type Cell struct {
	cellType CellType
	team     string
}

type Grid struct {
	cells [][]Cell
}

func NewEmptyGrid(x, y int) (grid *Grid) {
	cells := make([][]Cell, y)
	for i := range cells {
		cells[i] = make([]Cell, x)
	}

	grid = &Grid{cells}
	return
}

func NewGridFromCells(cells [][]Cell) (grid *Grid) {
	grid = &Grid{cells}
	return
}

func NewGridFromAscii(s []string) (grid *Grid, err error) {
	grid = NewEmptyGrid(len(s[0]), len(s))
	for i, row := range s {
		for j, b := range row {
			var cellType CellType
			if cellType, err = mapRuneToCellType(b); err != nil {
				return nil, err
			}
			grid.cells[i][j] = Cell{cellType, ""}
		}
	}
	return
}

func (grid Grid) Width() int {
	return len(grid.cells[0])
}

func (grid Grid) Height() int {
	return len(grid.cells)
}

func mapRuneToCellType(r rune) (CellType, error) {
	cellTypeMap := map[rune]CellType{
		'.': nilCell,
		'a': aCell,
		'b': bCell,
		'c': cCell,
		'W': wallCell,
	}
	if cellType, ok := cellTypeMap[r]; ok {
		return cellType, nil
	} else {
		return nilCell, errors.New("No cell type matches rune '" + string(r) + "'")
	}
}
