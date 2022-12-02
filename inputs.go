package aoc22

import (
	"github.com/bitfield/script"
	"strings"
)

func LoadInputs(file string) []string{
	contents, err := script.File(file).String()
  if err != nil {
    panic(err)
  }
  return strings.Split(contents, "\n")
}