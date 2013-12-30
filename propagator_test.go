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
	_ = p.Propagate(teamOneCell(aCell), "c3")
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
	_ = g.Set("c3", teamOneCell(aCell))
	p := NewPropagator(g)
	p.Propagate(teamOneCell(aCell), "c3")
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
	_ = g.Set("c3", teamOneCell(bCell))
	p := NewPropagator(g)
	p.Propagate(teamOneCell(bCell), "c3")
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
	p.Propagate(teamOneCell(aCell), "d4")
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
	p.Propagate(teamOneCell(bCell), "d5")
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

func TestSingleCCellPropagation(t *testing.T) {
	g := getSmallEmptyGrid()
	p := NewPropagator(g)
	p.Propagate(teamOneCell(cCell), "c3")
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

func TestDoubleCCellPropagation(t *testing.T) {
	g := getSmallEmptyGrid()
	_ = g.Set("c3", teamOneCell(cCell))
	p := NewPropagator(g)
	p.Propagate(teamOneCell(cCell), "c3")
	expectedCells := []string{
		".....",
		".....",
		"..w..",
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

func TestMultipleWallPropagation(t *testing.T) {
	cells := []string{
		"...c...",
		".abcba.",
		"..ccc..",
		".accca.",
		"..ccc..",
		".abcba.",
		"...cc..",
	}
	g, _ := NewGridFromAscii(cells)
	p := NewPropagator(g)
	p.Propagate(teamOneCell(cCell), "d7")
	expectedCells := []string{
		"...w...",
		".abwba.",
		"..www..",
		".awwwa.",
		"..www..",
		".abwba.",
		"...ww..",
	}
	expectedGrid, _ := NewGridFromAscii(expectedCells)
	if !reflect.DeepEqual(g, expectedGrid) {
		gridString, _ := StringForGrid(g)
		expectedGridString, _ := StringForGrid(expectedGrid)
		t.Errorf("Propagated grid is different than expected.\nExpected:\n%sOutput:\n%s\n", expectedGridString, gridString)
	}
}

func TestTwoTeamPropagation(t *testing.T) {
	cells := []string{
		"........",
		".abaABA.",
		"........",
	}
	g, _ := NewGridFromAscii(cells)
	p := NewPropagator(g)
	p.Propagate(teamOneCell(aCell), "d2")
	expectedCells := []string{
		"...aa...",
		".abbbBA.",
		"...aa...",
	}
	expectedGrid, _ := NewGridFromAscii(expectedCells)
	if !reflect.DeepEqual(g, expectedGrid) {
		gridString, _ := StringForGrid(g)
		expectedGridString, _ := StringForGrid(expectedGrid)
		t.Errorf("Propagated grid is different than expected.\nExpected:\n%sOutput:\n%s\n", expectedGridString, gridString)
	}
}

func TestTeamOneAndTeamNilPropagation(t *testing.T) {
	cells := []string{
		".......",
		"...x...",
		"...x...",
		"..x.x..",
		".......",
	}
	g, _ := NewGridFromAscii(cells)
	p := NewPropagator(g)
	p.Propagate(teamOneCell(bCell), "d3")
	expectedCells := []string{
		"...a...",
		"..aba..",
		"..aba..",
		"..xax..",
		".......",
	}
	expectedGrid, _ := NewGridFromAscii(expectedCells)
	if !reflect.DeepEqual(g, expectedGrid) {
		gridString, _ := StringForGrid(g)
		expectedGridString, _ := StringForGrid(expectedGrid)
		t.Errorf("Propagated grid is different than expected.\nExpected:\n%sOutput:\n%s\n", expectedGridString, gridString)
	}
}

func TestComplexFourTeamPropagation(t *testing.T) {
	cells := []string{
		"...wwWW...",
		"..a...#.q.",
		".abaxx#qrq",
		"..a..Q#.q.",
		".zy..qqqq.",
		"..z...AB..",
	}
	g, _ := NewGridFromAscii(cells)
	p := NewPropagator(g)
	p.Propagate(teamOneCell(cCell), "f4")
	expectedCells := []string{
		"...wwWW...",
		"..aaab#.q.",
		".abbbc#qrq",
		"..aaab#aba",
		".zy.abbbba",
		"..z..abBa.",
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
