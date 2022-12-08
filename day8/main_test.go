package main

import (
	"testing"
)

func TestInput(t *testing.T) {

	grid, xSize, ySize := LoadGrid("inputs_test.txt")

	if ok, _ := grid.IsVisible(Coord{1, 1}, xSize, ySize); !ok {
		t.Fail()
	}

	if ok, _ := grid.IsVisible(Coord{1, 2}, xSize, ySize); !ok {
		t.Fail()
	}

	if ok, _ := grid.IsVisible(Coord{1, 3}, xSize, ySize); ok {
		t.Fail()
	}
	if ok, _ := grid.IsVisible(Coord{1, 3}, xSize, ySize); ok {
		t.Fail()
	}

	if ok, _ := grid.IsVisible(Coord{2, 1}, xSize, ySize); !ok {
		t.Fail()
	}
}

func TestInput1(t *testing.T) {
	grid, xSize, ySize := LoadGrid("inputs_test1.txt")
	if ok, _ := grid.IsVisible(Coord{2, 2}, xSize, ySize); ok {
		t.Fail()
	}
}

func TestInput2(t *testing.T) {
	grid, xSize, ySize := LoadGrid("inputs_test.txt")
	bestVisiblityScore := 0
	for k, _ := range grid {
		if ok, score := grid.IsVisible(k, xSize, ySize); ok {
			if score > bestVisiblityScore {
				bestVisiblityScore = score
			}
		}
	}
	if bestVisiblityScore != 8 {
		t.Log(bestVisiblityScore)
		t.Fail()
	}
}
