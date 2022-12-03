package main

import (
  "github.com/lukegriffith/aoc22"
  "fmt"
)

type void struct{}

type item rune

// :-) in golang A is before a.
func (i item) getValue() int {
  val := int(i) - 96
  if (val < 0) {
    val = val + 58
  }
  return val
}

var (
  prioritySum = 0
  contentTypeMap map[item]void
  itemScore map[item]int = map[item]int{
    'p': 16,
    'L': 38,
    'P': 42,
    'v': 22,
    't': 20,
    's': 19,
  }
)

func splitBackpack(s string) (string, string) {
  slen := len(s)/2
  return s[0:slen], s[slen:]
}


func main() { 
  input := aoc22.LoadInputs("input.txt")
  for l, i := range input {
    fmt.Println(l)
    found := false
    contentTypeMap = make(map[item]void)
    s1, s2 := splitBackpack(i)
    fmt.Println(s1," | ",s2)
    for _, c := range s1 {
      contentTypeMap[item(c)] = void{}
    }
    for _, c := range s2 {
      if _, ok := contentTypeMap[item(c)]; ok {
        fmt.Println("item", string(c), "in both.")
        if found {
          fmt.Println("Theres more than one!")
          continue
        }
        found = true
        prioritySum = prioritySum + item(c).getValue()
      }
    }
    if ! found {
      fmt.Println("item not found.")
    }

  }
  fmt.Println(prioritySum)

}
