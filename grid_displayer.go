package main

import (
	"bytes"
	"fmt"
)

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
