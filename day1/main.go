package main

import (
  "github.com/bitfield/script"
  "fmt"
  "strings"
  "strconv"
  "sort"
)

var (
  input string           = "day1/input.txt"
  elfCounter int         = 1
  elfCalories            = 0
  calories               = []int{}

)

func main() {
  contents, err := script.File(input).String()
  if err != nil {
    panic(err)
  }
  inputs := strings.Split(contents, "\n")

  for _, i := range inputs {
    if i == "" {
      calories = append(calories, elfCalories)
      elfCalories = 0
      elfCounter++
    } else {
      intI, err := strconv.Atoi(i)
      if err != nil {
        panic(err)
      }
      elfCalories = elfCalories + intI
    }
  }
  sort.Slice(calories, func(i, j int) bool {
    return calories[i] < calories[j]
  })

  top3 := calories[len(calories)-3:]
  val1 := top3[0]
  val2 := top3[1]
  val3 := top3[2]
  fmt.Println(val1 + val2 + val3)
}
