package aoc22

import (
	"strings"

	"github.com/bitfield/script"
)

func LoadInputs(file string) []string {
	contents, err := script.File(file).String()
	if err != nil {
		panic(err)
	}
	return strings.Split(contents, "\n")
}

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}
