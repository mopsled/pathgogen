package main

import (
	"testing"
)

func TestCreateGridWithAscii(t *testing.T) {
	cells := []string{
		"a...a",
		".b.b.",
		"..c..",
	}
	g, err := NewGridFromAscii(cells)
	if err != nil {
		t.Errorf("Grid created from ascii produced an error: %s", err)
	}
	if g.Width() != 5 {
		t.Error("Grid has the wrong width for its initializing string")
	}
	if g.Height() != 3 {
		t.Error("Grid has the wrong width for its initializing string")
	}
	if g.cells[0][0].cellType != aCell {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if g.cells[2][2].cellType != cCell {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if g.cells[0][2].cellType != nilCell {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if g.cells[1][3].cellType != bCell {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
}

func TestStringForGridForTeam1(t *testing.T) {
	cells := []string{
		".abcw.",
	}
	g, err := NewGridFromAscii(cells)
	if err != nil {
		t.Errorf("Grid created from ascii for team 1 produced an error: %s", err)
	}
	if !(g.cells[0][1].cellType == aCell && g.cells[0][1].team == 1) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][2].cellType == bCell && g.cells[0][2].team == 1) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][3].cellType == cCell && g.cells[0][3].team == 1) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][4].cellType == wallCell && g.cells[0][4].team == 1) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
}

func TestStringForGridForTeam2(t *testing.T) {
	cells := []string{
		".ABCW.",
	}
	g, err := NewGridFromAscii(cells)
	if err != nil {
		t.Errorf("Grid created from ascii for team 1 produced an error: %s", err)
	}
	if !(g.cells[0][1].cellType == aCell && g.cells[0][1].team == 2) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][2].cellType == bCell && g.cells[0][2].team == 2) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][3].cellType == cCell && g.cells[0][3].team == 2) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][4].cellType == wallCell && g.cells[0][4].team == 2) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
}

func TestStringForGridForTeam3(t *testing.T) {
	cells := []string{
		".qrst.",
	}
	g, err := NewGridFromAscii(cells)
	if err != nil {
		t.Errorf("Grid created from ascii for team 1 produced an error: %s", err)
	}
	if !(g.cells[0][1].cellType == aCell && g.cells[0][1].team == 3) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][2].cellType == bCell && g.cells[0][2].team == 3) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][3].cellType == cCell && g.cells[0][3].team == 3) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][4].cellType == wallCell && g.cells[0][4].team == 3) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
}

func TestStringForGridForTeam4(t *testing.T) {
	cells := []string{
		".QRST.",
	}
	g, err := NewGridFromAscii(cells)
	if err != nil {
		t.Errorf("Grid created from ascii for team 1 produced an error: %s", err)
	}
	if !(g.cells[0][1].cellType == aCell && g.cells[0][1].team == 4) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][2].cellType == bCell && g.cells[0][2].team == 4) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][3].cellType == cCell && g.cells[0][3].team == 4) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][4].cellType == wallCell && g.cells[0][4].team == 4) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
}

func TestStringForGridForTeam0(t *testing.T) {
	cells := []string{
		".xyz#.",
	}
	g, err := NewGridFromAscii(cells)
	if err != nil {
		t.Errorf("Grid created from ascii for team 1 produced an error: %s", err)
	}
	if !(g.cells[0][0].cellType == nilCell && g.cells[0][0].team == 0) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][1].cellType == aCell && g.cells[0][1].team == 0) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][2].cellType == bCell && g.cells[0][2].team == 0) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][3].cellType == cCell && g.cells[0][3].team == 0) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
	if !(g.cells[0][4].cellType == wallCell && g.cells[0][4].team == 0) {
		t.Error("Unexpected cellType in grid initialized with ascii")
	}
}

func TestStringForGridWithASimpleGrid(t *testing.T) {
	cells := []string{
		".....",
		".....",
		"..a..",
		".....",
		".....",
	}
	g, err := NewGridFromAscii(cells)
	if err != nil {
		t.Errorf("Grid created from ascii produced an error: %s", err)
	}
	s, err := StringForGrid(g)
	if err != nil {
		t.Errorf("Error creating a string from grid: %s", err)
	}
	expectedString := ".....\n.....\n..a..\n.....\n.....\n"
	if s != expectedString {
		t.Errorf("Unexpected output for grid. Expected:\n%sGot:\n%s", expectedString, s)
	}
}

func TestCreateGridWithAsciiWithTwoTeams(t *testing.T) {
	cells := []string{
		"..a.....C",
		"b..wwC...",
		"..aWW.a..",
		".Cbwwaba.",
		".bcbBBbw.",
		"..b..wWwc",
	}

	g, err := NewGridFromAscii(cells)
	if err != nil {
		t.Errorf("Grid created from ascii produced an error: %s", err)
	}
	s, err := StringForGrid(g)
	if err != nil {
		t.Errorf("Error creating a string from grid: %s", err)
	}
	expectedString := "..a.....C\nb..wwC...\n..aWW.a..\n.Cbwwaba.\n.bcbBBbw.\n..b..wWwc\n"
	if s != expectedString {
		t.Errorf("Unexpected output for grid. Expected:\n%sGot:\n%s", expectedString, s)
	}
}

func TestCreateGridWithAsciiWithAllTeams(t *testing.T) {
	cells := []string{
		"ax#W#xA",
		"zb.#.Bz",
		"sycTCyS",
		"qr.w.RQ",
	}

	g, err := NewGridFromAscii(cells)
	if err != nil {
		t.Errorf("Grid created from ascii produced an error: %s", err)
	}
	s, err := StringForGrid(g)
	if err != nil {
		t.Errorf("Error creating a string from grid: %s", err)
	}
	expectedString := "ax#W#xA\nzb.#.Bz\nsycTCyS\nqr.w.RQ\n"
	if s != expectedString {
		t.Errorf("Unexpected output for grid. Expected:\n%sGot:\n%s", expectedString, s)
	}
}
