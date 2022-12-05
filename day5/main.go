package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/lukegriffith/aoc22"
	"github.com/lukegriffith/aoc22/stack"
)

type crate rune

type move struct {
	source      int
	destination int
	quantity    int
	original    string
}

type crates struct {
	rows []stack.Stack[crate]
}

func LoadCrates(s []string) crates {
	rowString := s[len(s)-1]
	c := crates{[]stack.Stack[crate]{}}
	for _, i := range rowString {
		if i != ' ' {
			c.rows = append(c.rows, stack.Stack[crate]{})
		}
	}
	totalRows := len(c.rows)
	for i := len(s) - 2; i >= 0; i-- {
		emptiesSeen := 0
		crateRow := 0
		row := s[i]
		for _, col := range row {
			if unicode.IsUpper(col) {
				c.rows[crateRow].Push(crate(col))
				crateRow++
				emptiesSeen = 0
				if crateRow > totalRows {
					panic("Somethings gone wrong")
				}
			}
			if unicode.IsSpace(col) {
				emptiesSeen++
			}
			if emptiesSeen == 4 {
				emptiesSeen = 0
				crateRow++
			}
		}
	}
	return c
}

func LoadMoves(s []string) []move {
	moveArr := []move{}
	for _, i := range s {
		m := move{0, 0, 0, i}
		ii := strings.Split(i, " ")
		m.original = i
		strInt1, err := strconv.Atoi(string(ii[1]))
		aoc22.CheckErr(err)
		m.quantity = strInt1
		strInt2, err := strconv.Atoi(ii[3])
		aoc22.CheckErr(err)
		m.source = strInt2
		strInt3, err := strconv.Atoi(ii[5])
		aoc22.CheckErr(err)
		m.destination = strInt3

		moveArr = append(moveArr, m)
	}
	return moveArr
}

func main() {
	crateInput := aoc22.LoadInputs("crates.txt")
	input := aoc22.LoadInputs("input.txt")
	cratesData := LoadCrates(crateInput)
	moves := LoadMoves(input)
	fmt.Println("init", cratesData)
	for _, m := range moves {
		fmt.Println(m)
		movebuffer := stack.Stack[crate]{}
		for i := m.quantity - 1; i >= 0; i-- {
			fmt.Println("pre:", cratesData)
			c, ok := cratesData.rows[m.source-1].Pop()

			if !ok {
				panic("nothing left in source")
				//fmt.Println(cratesData)
				//fmt.Println("nothing left in source")
				//continue
			}
			//cratesData.rows[m.destination-1].Push(c)
			movebuffer.Push(c)

		}
		for !movebuffer.IsEmpty() {
			c, _ := movebuffer.Pop()
			cratesData.rows[m.destination-1].Push(c)
		}
		fmt.Println("post:", cratesData)
	}

	message := ""
	for _, c := range cratesData.rows {
		r, ok := c.Pop()
		if !ok {
			fmt.Println("warn: nothing in row for message.")
		}
		message = message + string(r)
	}
	fmt.Println(cratesData)
	fmt.Println(message)
}
