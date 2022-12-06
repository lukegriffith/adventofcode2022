package main

import (
	"fmt"

	"github.com/lukegriffith/aoc22"
	"github.com/lukegriffith/aoc22/queue"
)

func decodeMarker(input string, markerLength int) int {
	q := queue.NewQueueCap[rune](4)

	for _, i := range input {
		q.Push(i)
	}
	type void struct{}
	for q.Position()+markerLength < q.Length() {
		runeMap := map[rune]void{}
		for i := 0; i < markerLength; i++ {
			runeMap[q.PeekMore(i)] = void{}
		}

		keys := 0
		for _, _ = range runeMap {
			keys++
		}
		if keys == markerLength {
			return q.Position() + markerLength
		}
		q.Next()

	}
	return 0
}

func main() {
	input := aoc22.LoadInputs("input.txt")[0]
	fmt.Println("part 1")
	i := decodeMarker(input, 4)
	fmt.Println(i)
	fmt.Println("part 2")
	i = decodeMarker(input, 14)
	fmt.Println(i)
}
