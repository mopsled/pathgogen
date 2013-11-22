package main

type CellType int

const (
	nilCell CellType = iota
	aCell
	bCell
	cCell
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
