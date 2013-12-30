package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func NewGridFromAscii(s []string) (grid *Grid, err error) {
	grid = NewEmptyGrid(len(s[0]), len(s))
	for i, row := range s {
		for j, b := range row {
			var cellType CellType
			if cellType, err = mapRuneToCellType(b); err != nil {
				return nil, err
			}
			var team int
			if team, err = mapRuneToTeam(b); err != nil {
				return nil, err
			}
			grid.cells[i][j] = Cell{cellType, team}
		}
	}
	return
}

func StringForGrid(grid *Grid) (string, error) {
	var output bytes.Buffer
	for _, row := range grid.cells {
		for _, c := range row {
			runeForCell, err := mapCellToRune(c)
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
	cellRegexes := []string{"\\.", "[aAqQx]", "[bBrRy]", "[cCsSz]", "[wWtT#]"}
	cellTypes := []CellType{nilCell, aCell, bCell, cCell, wallCell}

	var cellType CellType = -1
	for i, cellRegex := range cellRegexes {
		if matched, _ := regexp.MatchString(cellRegex, string(r)); matched {
			cellType = cellTypes[i]
		}
	}
	if cellType != -1 {
		return cellType, nil
	} else {
		return nilCell, fmt.Errorf("No cell type matches rune '%s'", string(r))
	}
}

func mapRuneToTeam(r rune) (int, error) {
	teamRegexes := []string{"[xyz#.]", "[abcw]", "[ABCW]", "[qrst]", "[QRST]"}

	team := -1
	for i, teamRegex := range teamRegexes {
		if matched, _ := regexp.MatchString(teamRegex, string(r)); matched {
			team = i
		}
	}
	if team != -1 {
		return team, nil
	} else {
		return team, fmt.Errorf("No team matches rune '%s'", string(r))
	}
}

func mapCellToRune(cell Cell) (rune, error) {
	teams := [][]rune{[]rune(".xyz#"), []rune(".abcw"), []rune(".ABCW"), []rune(".qrst"), []rune(".QRST")}

	if cell.team < 0 || cell.team > 4 {
		return '?', fmt.Errorf("No rune for cell types on team %d", cell.team)
	}

	if cell.cellType == nilCell && cell.team != 0 {
		return '?', fmt.Errorf("Not possible to create a nilCell on non-zero team (%d)", cell.team)
	}

	return teams[cell.team][cell.cellType], nil
}
