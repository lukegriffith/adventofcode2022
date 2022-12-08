package main

import (
	"fmt"

	"github.com/lukegriffith/aoc22"
)

type CoordMap map[Coord]int

func (c CoordMap) IsVisible(coord Coord, xSize int, ySize int) (bool, int) {

	x := coord.X
	y := coord.Y
	value := c[coord]

	visibleL, visibleR, visibleT, visibleB := true, true, true, true
	totalViewScore := 0
	score0 := 0
	for i := x - 1; i >= 0; i-- {
		score0++
		eval_coodr := c[Coord{i, y}]
		if eval_coodr >= value {
			visibleL = false
			break
		}
	}
	score1 := 0
	for i := x + 1; i < xSize; i++ {
		score1++
		eval_coodr := c[Coord{i, y}]
		if eval_coodr >= value {
			visibleR = false
			break
		}
	}
	score2 := 0
	for i := y - 1; i >= 0; i-- {
		score2++
		eval_coodr := c[Coord{x, i}]
		if eval_coodr >= value {
			visibleT = false
			break
		}
	}
	score3 := 0
	for i := y + 1; i < ySize; i++ {
		score3++
		eval_coodr := c[Coord{x, i}]
		if eval_coodr >= value {
			visibleB = false
			break
		}
	}
	totalViewScore = score0 * score1 * score2 * score3
	return visibleB || visibleL || visibleR || visibleT, totalViewScore
}

type Coord struct {
	X int
	Y int
}

func LoadGrid(s string) (CoordMap, int, int) {
	inputs := aoc22.LoadInputs(s)
	grid := make(CoordMap)
	for i, ival := range inputs {
		for x, xval := range ival {
			intX := int(xval) - '0'
			grid[Coord{x, i}] = intX
		}
	}
	return grid, len(inputs[0]), len(inputs)
}

func main() {
	grid, xSize, ySize := LoadGrid("inputs.txt")
	visible := 0
	bestVisiblityScore := 0
	for k, _ := range grid {
		fmt.Println(k)
		if ok, score := grid.IsVisible(k, xSize, ySize); ok {
			if score > bestVisiblityScore {
				bestVisiblityScore = score
			}
			visible++
			fmt.Println("visible")
		} else {
			fmt.Println("not visible")
		}
	}
	fmt.Println(visible, bestVisiblityScore)
}
