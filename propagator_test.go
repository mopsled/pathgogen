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

func TestACellPropagation(t *testing.T) {
	g := getSmallEmptyGrid()
	p := NewPropagator(g)
	p.Propagate(teamlessCell(aCell), "c3")
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
		t.Errorf("Propagated grid is different than expected.\nExpected:\n%sOutput:\n%s\n", gridString, expectedGridString)
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
