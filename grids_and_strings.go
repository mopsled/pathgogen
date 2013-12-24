package main

import (
	"bytes"
	"fmt"
)

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

func StringForGrid(grid *Grid) (string, error) {
	var output bytes.Buffer
	for _, row := range grid.cells {
		for _, c := range row {
			runeForCell, err := mapCellTypeToRune(c.cellType)
			if err != nil {
				return "", fmt.Errorf("Could not map given grid to a string: %s", err)
			}
			output.WriteRune(runeForCell)
		}
		output.WriteRune('\n')
	}
	return output.String(), nil
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

func mapCellTypeToRune(cellType CellType) (rune, error) {
	runeToCellTypeMap := map[CellType]rune{
		nilCell:  '.',
		aCell:    'a',
		bCell:    'b',
		cCell:    'c',
		wallCell: 'w',
	}
	if runeForCellType, ok := runeToCellTypeMap[cellType]; ok {
		return runeForCellType, nil
	} else {
		return '?', fmt.Errorf("No rune for cell type '%s'", cellType)
	}
}
