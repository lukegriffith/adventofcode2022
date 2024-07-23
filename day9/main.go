package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/lukegriffith/aoc22"
)

type CoordMap map[Coord]int

type Coord struct {
	X int
	Y int
}

type void struct{}

type Entity struct {
	CurrentPosition  Coord
	PreviousPosition map[Coord]void
}

func NewEntity() *Entity {
	e := Entity{Coord{0, 0}, make(map[Coord]void)}
	e.PreviousPosition[Coord{0, 0}] = void{}
	return &e
}

func (e *Entity) Move(direction rune) {
	switch direction {
	case 'U':
		e.CurrentPosition = Coord{e.CurrentPosition.X, e.CurrentPosition.Y + 1}
		e.PreviousPosition[e.CurrentPosition] = void{}
	case 'D':
		e.CurrentPosition = Coord{e.CurrentPosition.X, e.CurrentPosition.Y - 1}
		e.PreviousPosition[e.CurrentPosition] = void{}
	case 'L':
		e.CurrentPosition = Coord{e.CurrentPosition.X - 1, e.CurrentPosition.Y}
		e.PreviousPosition[e.CurrentPosition] = void{}
	case 'R':
		e.CurrentPosition = Coord{e.CurrentPosition.X + 1, e.CurrentPosition.Y}
		e.PreviousPosition[e.CurrentPosition] = void{}
	}
}

func (e *Entity) MoveToEntity(target *Entity) {
	if math.Abs(float64(e.CurrentPosition.X-target.CurrentPosition.X)) < 2 &&
		math.Abs(float64(e.CurrentPosition.Y-target.CurrentPosition.Y)) < 2 {
		return
	}
	c := e.CurrentPosition
	if target.CurrentPosition.X > e.CurrentPosition.X {
		c.X = e.CurrentPosition.X + 1
	}
	if target.CurrentPosition.X < e.CurrentPosition.X {
		c.X = e.CurrentPosition.X - 1
	}
	if target.CurrentPosition.Y > e.CurrentPosition.Y {
		c.Y = e.CurrentPosition.Y + 1
	}
	if target.CurrentPosition.Y < e.CurrentPosition.Y {
		c.Y = e.CurrentPosition.Y - 1
	}

	e.CurrentPosition = c
	e.PreviousPosition[c] = void{}

}

/*

0,0  0,0    0,0  abs(c-t) > 1 move 1
0,1  0,0    0,0
0,2  0,1    0,1
1,2  0,2
2,3  1,3

*/

func day9(inputs []string) (int, int) {

	head := NewEntity()
	tail := NewEntity()

	knots := []*Entity{}

	for i := 0; i < 10; i++ {
		knots = append(knots, NewEntity())
	}

	for _, v := range inputs {
		splitInputs := strings.Split(v, " ")
		dir := splitInputs[0]
		steps := splitInputs[1]
		stepsI, err := strconv.Atoi(steps)
		aoc22.CheckErr(err)
		for i := 0; i < stepsI; i++ {
			d := rune(dir[0])
			// part 1
			head.Move(d)
			tail.MoveToEntity(head)
			//part 2
			knots[0].Move(d)
			last := knots[0]
			for _, k := range knots[1:] {
				k.MoveToEntity(last)
				last = k
			}
		}
	}

	positions := 0
	for _, _ = range tail.PreviousPosition {
		positions++
	}

	tailPosition := 0
	for _, _ = range knots[len(knots)-1].PreviousPosition {
		tailPosition++
	}
	return positions, tailPosition
}

func main() {
	inputs := aoc22.LoadInputs("input.txt")
	fmt.Println(day9(inputs))
}

// 0,0 0,0
// 1,1 0,0
