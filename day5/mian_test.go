package main

import (
	"testing"

	"github.com/lukegriffith/aoc22"
)

func TestLoadCrates(t *testing.T) {
	crates := aoc22.LoadInputs("test_crates.txt")
	c := LoadCrates(crates)
	if c.rows[0].Length() != 2 {
		t.Fail()
	}
	if c.rows[1].Length() != 2 {
		t.Fail()
	}
	if c.rows[2].Length() != 3 {
		t.Fail()
	}
	r, _ := c.rows[0].Pop()
	if r != 'A' {
		t.Fail()
	}
	r, _ = c.rows[0].Pop()
	if r != 'C' {
		t.Fail()
	}
	r, _ = c.rows[1].Pop()
	if r != 'B' {
		t.Fail()
	}
	r, _ = c.rows[1].Pop()
	if r != 'D' {
		t.Fail()
	}
}

func TestLoadMoves(t *testing.T) {
	m := aoc22.LoadInputs("test_input.txt")
	m1 := LoadMoves(m)

	if m1[0].quantity != 3 {
		t.Fail()
	}
	if m1[0].source != 2 {
		t.Fail()
	}
	if m1[0].destination != 5 {
		t.Fail()
	}
	if m1[1].quantity != 20 {
		t.Fail()
	}
	if m1[1].source != 9 {
		t.Fail()
	}
	if m1[1].destination != 6 {
		t.Fail()
	}
}
