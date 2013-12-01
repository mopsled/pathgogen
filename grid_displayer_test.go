package main

import (
	"testing"
)

func TestStringForGridWithASimpleGrid(t *testing.T) {
	cells := []string{
		".....",
		".....",
		"..a..",
		".....",
		".....",
	}
	g, _ := NewGridFromAscii(cells)
	s, err := StringForGrid(g)
	if err != nil {
		t.Errorf("Error creating a string from grid: %s", err)
	}
	expectedString := ".....\n.....\n..a..\n.....\n.....\n"
	if s != expectedString {
		t.Errorf("Unexpected output for grid. Expected:\n%sGot:\n%s", expectedString, s)
	}
}

func TestStringForGridWithAModeratelyComplexGrid(t *testing.T) {
	cells := []string{
		"..a......",
		"b..ww....",
		"..aww.a..",
		"..bwwaba.",
		".bcb..b..",
		"..b......",
	}

	g, _ := NewGridFromAscii(cells)
	s, err := StringForGrid(g)
	if err != nil {
		t.Errorf("Error creating a string from grid: %s", err)
	}
	expectedString := "..a......\nb..ww....\n..aww.a..\n..bwwaba.\n.bcb..b..\n..b......\n"
	if s != expectedString {
		t.Errorf("Unexpected output for grid. Expected:\n%sGot:\n%s", expectedString, s)
	}
}
