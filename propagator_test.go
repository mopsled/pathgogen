package main

import (
	"reflect"
	"testing"
)

func TestCreatePropagator(t *testing.T) {
	g := getSmallEmptyGrid()
	p := NewPropagator(g)
	if p == nil {
		t.Error("Can't create a propagator with a grid")
	}
}

func TestACellOnEmptyBoardPropagation(t *testing.T) {
	g := getSmallEmptyGrid()
	p := NewPropagator(g)
	_ = p.Propagate(teamlessCell(aCell), "c3")
	expectedCells := []string{
		".....",
		".....",
		"..a..",
		".....",
		".....",
	}
	expectedGrid, _ := NewGridFromAscii(expectedCells)
	if !reflect.DeepEqual(g, expectedGrid) {
		gridString, _ := StringForGrid(g)
		expectedGridString, _ := StringForGrid(expectedGrid)
		t.Errorf("Propagated grid is different than expected.\nExpected:\n%sOutput:\n%s\n", expectedGridString, gridString)
	}
}

func TestDoubleACellPropagation(t *testing.T) {
	g := getSmallEmptyGrid()
	_ = g.Set("c3", teamlessCell(aCell))
	p := NewPropagator(g)
	p.Propagate(teamlessCell(aCell), "c3")
	expectedCells := []string{
		".....",
		"..a..",
		".aba.",
		"..a..",
		".....",
	}
	expectedGrid, _ := NewGridFromAscii(expectedCells)
	if !reflect.DeepEqual(g, expectedGrid) {
		gridString, _ := StringForGrid(g)
		expectedGridString, _ := StringForGrid(expectedGrid)
		t.Errorf("Propagated grid is different than expected.\nExpected:\n%sOutput:\n%s\n", expectedGridString, gridString)
	}
}

func TestDoubleBCellPropagation(t *testing.T) {
	g := getSmallEmptyGrid()
	_ = g.Set("c3", teamlessCell(bCell))
	p := NewPropagator(g)
	p.Propagate(teamlessCell(bCell), "c3")
	expectedCells := []string{
		"..a..",
		".aba.",
		"abcba",
		".aba.",
		"..a..",
	}
	expectedGrid, _ := NewGridFromAscii(expectedCells)
	if !reflect.DeepEqual(g, expectedGrid) {
		gridString, _ := StringForGrid(g)
		expectedGridString, _ := StringForGrid(expectedGrid)
		t.Errorf("Propagated grid is different than expected.\nExpected:\n%sOutput:\n%s\n", expectedGridString, gridString)
	}
}

func TestDoubleAAndBCellPropagation(t *testing.T) {
	cells := []string{
		".......",
		".......",
		"..aaa..",
		"..bab..",
		"..aaa..",
		".......",
		".......",
	}
	g, _ := NewGridFromAscii(cells)
	p := NewPropagator(g)
	p.Propagate(teamlessCell(aCell), "d4")
	expectedCells := []string{
		".......",
		"..aaa..",
		".abbba.",
		"..bbb..",
		".abbba.",
		"..aaa..",
		".......",
	}
	expectedGrid, _ := NewGridFromAscii(expectedCells)
	if !reflect.DeepEqual(g, expectedGrid) {
		gridString, _ := StringForGrid(g)
		expectedGridString, _ := StringForGrid(expectedGrid)
		t.Errorf("Propagated grid is different than expected.\nExpected:\n%sOutput:\n%s\n", expectedGridString, gridString)
	}
}

func TestBNextToAAndBCellPropagation(t *testing.T) {
	cells := []string{
		".......",
		".aaaaa.",
		".a...a.",
		".a.b.a.",
		".a...a.",
		".aaaaa.",
		".......",
	}
	g, _ := NewGridFromAscii(cells)
	p := NewPropagator(g)
	p.Propagate(teamlessCell(bCell), "d5")
	expectedCells := []string{
		".aaaaa.",
		"abbbbba",
		"abababa",
		"abababa",
		"abaaaba",
		"abbbbba",
		".aaaaa.",
	}
	expectedGrid, _ := NewGridFromAscii(expectedCells)
	if !reflect.DeepEqual(g, expectedGrid) {
		gridString, _ := StringForGrid(g)
		expectedGridString, _ := StringForGrid(expectedGrid)
		t.Errorf("Propagated grid is different than expected.\nExpected:\n%sOutput:\n%s\n", expectedGridString, gridString)
	}
}

func getSmallEmptyGrid() *Grid {
	cells := []string{
		".....",
		".....",
		".....",
		".....",
		".....",
	}
	g, _ := NewGridFromAscii(cells)
	return g
}
