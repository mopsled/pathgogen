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

func NewEmptyGrid(x, y int) *Grid {
	cells := make([][]Cell, y)
	for i := range cells {
		cells[i] = make([]Cell, x)
	}
	return &Grid{cells}
}

func NewGridFromCells(cells [][]Cell) *Grid {
	return &Grid{cells}
}

func (grid Grid) Width() int {
	return len(grid.cells[0])
}

func (grid Grid) Height() int {
	return len(grid.cells)
}

func (grid Grid) At(coordinate string) (*Cell, error) {
	row, column, err := RowAndColumnForCoordinate(coordinate, grid)
	if err != nil {
		return nil, err
	}
	return &grid.cells[row][column], nil
}

func (grid Grid) Set(coordinate string, cell *Cell) (err error) {
	row, column, err := RowAndColumnForCoordinate(coordinate, grid)
	if err != nil {
		return err
	}
	grid.cells[row][column] = *cell
	return
}

func RowAndColumnForCoordinate(coordinate string, grid Grid) (row, column int, err error) {
	if matched, _ := regexp.MatchString("[a-zA-Z]+[0-9]+", coordinate); !matched {
		err = fmt.Errorf("Unable to use strange coordinate '%s'. Try something like 'a5' or 'f27'.", coordinate)
		return
	}

	matches := regexp.MustCompile("([a-zA-Z]+)([0-9]+)").FindStringSubmatch(coordinate)
	letterString, numberString := matches[1], matches[2]
	number64, _ := strconv.ParseInt(numberString, 10, 0)
	number := int(number64)
	letterValue := letterIndexToInt(letterString)
	if number == 0 {
		err = fmt.Errorf("Grid coordinates are 1-based, not 0-based. Coordinate '%s' doesn't exist on the board.")
		return
	}
	if number > grid.Height() {
		err = fmt.Errorf("Coordinate '%s' doesn't exist on this grid, which has a height of %d", coordinate, grid.Height())
		return
	}
	if letterValue > grid.Width() {
		err = fmt.Errorf("Coordinate '%s' doesn't exist on this grid, which has a width of %d", coordinate, grid.Width())
		return
	}

	return (grid.Height() - number), (letterValue - 1), nil
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
