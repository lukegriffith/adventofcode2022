package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lukegriffith/aoc22"
)

func getRange(s string) []int {
	ss := strings.Split(s, "-")
	slow, err := strconv.Atoi(ss[0])
	aoc22.CheckErr(err)
	shigh, err := strconv.Atoi(ss[1])
	aoc22.CheckErr(err)
	var intRange []int

	for i := slow; i <= shigh; i++ {
		intRange = append(intRange, i)
	}
	return intRange
}

func elfRanges(s string) ([]int, []int) {
	ss := strings.Split(s, ",")
	s1 := ss[0]
	s2 := ss[1]

	range1 := getRange(s1)
	range2 := getRange(s2)
	return range1, range2

}

func main() {
	input := aoc22.LoadInputs("input.txt")
	rangesContained := 0

	for _, i := range input {
		e1range, e2range := elfRanges(i)

		e1low := e1range[0]
		e1high := e1range[len(e1range)-1]
		e2low := e2range[0]
		e2high := e2range[len(e2range)-1]

		if e1low <= e2low && e1high >= e2high {
			rangesContained = rangesContained + 1
			continue
		}
		if e2low <= e1low && e2high >= e1high {
			rangesContained = rangesContained + 1
			continue
		}

		type void struct{}
		var rangeMap = make(map[int]void)

		for _, i := range e1range {
			rangeMap[i] = void{}
		}

		for _, i := range e2range {
			if _, ok := rangeMap[i]; ok {
				rangesContained = rangesContained + 1
				break
			}
		}

	}
	fmt.Println(rangesContained)
}
