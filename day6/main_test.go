package main

import "testing"

// Part 1
func TestDecodePacket(t *testing.T) {

	tests := map[string]int{
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	}

	for k, v := range tests {
		i := decodeMarker(k, 4)
		if i != v {
			t.Log(k, "expected", v, "result", i)
			t.Fail()
		}
	}
}

// Part 2
func TestDecodeMessage(t *testing.T) {

	tests := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
		"nppdvjthqldpwncqszvftbrmjlhg":      23,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
	}

	for k, v := range tests {
		i := decodeMarker(k, 14)
		if i != v {
			t.Log(k, "expected", v, "result", i)
			t.Fail()
		}
	}
}
