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
  part1 bool = false
  part2 bool = true
)

func splitBackpack(s string) (string, string) {
  slen := len(s)/2
  return s[0:slen], s[slen:]
}


func main() { 
  input := aoc22.LoadInputs("input.txt")
  if part1 {
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

  if part2 {
    count := 0 
    badgeMap := make(map[item]int)
    elfBadgeCounter := make(map[item]void)
    badgePriority := 0 

    for _, i := range input { 
      for _, c := range i {
        if _, ok := elfBadgeCounter[item(c)]; !ok {
          elfBadgeCounter[item(c)] = void{}
        }
      }
      for k, _ := range elfBadgeCounter {
        if _, ok := badgeMap[item(k)]; ok {
          badgeMap[item(k)] = badgeMap[item(k)] + 1
        } else {
          badgeMap[item(k)] = 1
        }
      }
      fmt.Println(i)
      count++
      if count % 3 == 0 {
        fmt.Println("Break")
        for k, v := range badgeMap {
          if v == 3 {

            fmt.Println(count, string(k), v)
            badgePriority = badgePriority + k.getValue()
          }
        }
        fmt.Println(badgeMap)
        badgeMap = make(map[item]int)
      }

      
      //fmt.Println(elfBadgeCounter)
      elfBadgeCounter = make(map[item]void)
    }
    fmt.Println(badgePriority)
  }
}