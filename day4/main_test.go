package main

import "testing"

func TestGetRange(t *testing.T) {
	r := getRange("1-10")
	if len(r) != 10 {
		t.Fail()
	}

	if r[0] != 1 {
		t.Fail()
	}

	if r[9] != 10 {
		t.Fail()
	}
}

//   |     |
// |     |
//
// |     |
//     |    |
