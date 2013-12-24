package main

import (
	"fmt"
)

type Propagator struct {
	grid *Grid
}

type rowAndColumn struct {
	row, column int
}

func NewPropagator(grid *Grid) *Propagator {
	return &Propagator{grid}
}

func (p Propagator) Propagate(cell *Cell, coordinate string) error {
	row, column, err := RowAndColumnForCoordinate(coordinate, *p.grid)
	if err != nil {
		return err
	}
	return p.innerPropagate(cell, row, column)
}

func (p Propagator) innerPropagate(cell *Cell, row, column int) error {
	if cell.cellType == nilCell {
		return nil
	}
	currentCell := (*p.grid).cells[row][column]
	if currentCell.cellType > cell.cellType {
		return fmt.Errorf("Error propagating cell at (%d, %d). Trying to propagate with lesser cell type.", row, column)
	} else if currentCell.cellType == cell.cellType {
		(*p.grid).cells[row][column].cellType = upgrade(currentCell.cellType)
	} else {
		(*p.grid).cells[row][column].cellType = cell.cellType
	}
	neighbors := getNeighborCoordinates(row, column, *p.grid)
	for _, neighbor := range neighbors {
		neighborCell := (*p.grid).cells[neighbor.row][neighbor.column]
		if cell.greaterThan(neighborCell) {
			err := p.innerPropagate(&Cell{downgrade((*p.grid).cells[row][column].cellType), ""}, neighbor.row, neighbor.column)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (cell Cell) greaterThan(otherCell Cell) bool {
	return cell.cellType > otherCell.cellType
}

func downgrade(cellType CellType) CellType {
	return (CellType)(cellType - 1)
}

func upgrade(cellType CellType) CellType {
	return (CellType)(cellType + 1)
}

func getNeighborCoordinates(row, column int, grid Grid) []rowAndColumn {
	allNeighbors := []rowAndColumn{
		rowAndColumn{row + 1, column},
		rowAndColumn{row - 1, column},
		rowAndColumn{row, column + 1},
		rowAndColumn{row, column - 1},
	}
	validNeighbors := make([]rowAndColumn, 0)
	for _, neighbor := range allNeighbors {
		neighborHasAValidRow := neighbor.row >= 0 && neighbor.row < grid.Height()
		neighborHasAValidColumn := neighbor.column >= 0 && neighbor.column < grid.Width()
		if neighborHasAValidRow && neighborHasAValidColumn {
			validNeighbors = append(validNeighbors, neighbor)
		}
	}
	return validNeighbors
}
