package main

import (
	"testing"

	"github.com/lukegriffith/aoc22"
)

func TestInput(t *testing.T) {
	inputs := aoc22.LoadInputs("input_test.txt")
	value := day10(inputs)

	if value != 13140 {
		t.Log(value)
		t.Fail()
	}
}
