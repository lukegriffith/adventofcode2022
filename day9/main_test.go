package main

import (
	"testing"

	"github.com/lukegriffith/aoc22"
)

func TestMoveEntity(t *testing.T) {
	e := NewEntity()

	e.Move('R')
	e.Move('R')
	e.Move('R')

	keyCount := 0

	for _, _ = range e.PreviousPosition {
		keyCount++
	}

	if keyCount != 2 {
		t.Log(keyCount)
		t.Fail()
	}
}

func TestMoveTail(t *testing.T) {
	e1 := NewEntity()
	e2 := NewEntity()

	e1.Move('U')
	e2.MoveToEntity(e1)

	if e2.CurrentPosition.X != 0 && e2.CurrentPosition.Y != 0 {
		t.Fail()
	}

	e1.Move('U')
	e2.MoveToEntity(e1)
	if e2.CurrentPosition.X != 0 && e2.CurrentPosition.Y != 0 {
		t.Fail()
	}

	t.Log(e2)
}

func TestInput(t *testing.T) {
	inputs := aoc22.LoadInputs("input_test.txt")
	p, _ := day9(inputs)
	if p != 13 {
		t.Fail()
	}
}

func TestInputPart2(t *testing.T) {
	inputs := aoc22.LoadInputs("input_test2.txt")
	_, p := day9(inputs)
	if p != 36 {
		t.Log(p)
		t.Fail()
	}
}
