package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

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

func (grid Grid) At(coordinate string) (*Cell, error) {
	if matched, _ := regexp.MatchString("[a-zA-Z]+[0-9]+", coordinate); !matched {
		return nil, fmt.Errorf("Unable to use strange coordinate '%s'. Try something like 'a5' or 'f27'.", coordinate)
	}

	matches := regexp.MustCompile("([a-zA-Z]+)([0-9]+)").FindStringSubmatch(coordinate)
	letterString, numberString := matches[1], matches[2]
	number64, _ := strconv.ParseInt(numberString, 10, 0)
	number := int(number64)
	letterValue := letterIndexToInt(letterString)
	if number == 0 {
		return nil, fmt.Errorf("Grid coordinates are 1-based, not 0-based. Coordinate '%s' doesn't exist on the board.")
	}
	if number > grid.Height() {
		return nil, fmt.Errorf("Coordinate '%s' doesn't exist on this grid, which has a height of %d", coordinate, grid.Height())
	}
	if letterValue > grid.Width() {
		return nil, fmt.Errorf("Coordinate '%s' doesn't exist on this grid, which has a width of %d", coordinate, grid.Width())
	}

	return &grid.cells[grid.Height()-number][letterValue-1], nil
}

func mapRuneToCellType(r rune) (CellType, error) {
	cellTypeMap := map[rune]CellType{
		'.': nilCell,
		'a': aCell,
		'b': bCell,
		'c': cCell,
		'w': wallCell,
	}
	if cellType, ok := cellTypeMap[r]; ok {
		return cellType, nil
	} else {
		return nilCell, fmt.Errorf("No cell type matches rune '%s'", string(r))
	}
}

func letterIndexToInt(letters string) int {
	var letterValue, power int
	for _, letter := range reverse(letters) {
		value64, _ := strconv.ParseUint(string(letter), 36, 0)
		value := int(value64) - 9
		letterValue += value * int(math.Pow(26, float64(power)))
		power += 1
	}
	return letterValue
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
