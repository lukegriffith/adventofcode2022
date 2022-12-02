package main

import (
  "github.com/lukegriffith/aoc22"
  "strings"
  "fmt"
)

type move string

var (
  rockScore int = 1
  paperScore    = 2
  scissorScore  = 3
  lostScore     = 0
  drawScore     = 3
  winScore      = 6
  totalScore    = 0 
  moves         = map[string][]move{
    "rock":     []move{move("A"), move("X")},
    "paper":    []move{move("B"), move("Y")},
    "scissors": []move{move("C"), move("Z")},
  }
) 

func moveInSlice(a move, list []move) bool {
  for _, b := range list {
      if b == a {
          return true
      }
  }
  return false
}

func loadMove(moves string) (move, move) {
  m := strings.Split(moves, " ")
  return move(m[0]), move(m[1])
}

// 1 rock, 2, paper, 3 scissors
func idMove(m move) int {
  if moveInSlice(m, moves["rock"]) {
    return 0
  }
  if moveInSlice(m, moves["paper"]) {
    return 1
  }
  if moveInSlice(m, moves["scissors"]) {
    return 2
  }
  panic("Unable to id move")
}

// 1st return value is who won, second is score
func winningMovePair(m1 move, m2 move) (int, int) {
  var (
    score = 0
    moveMatrix = [3][3]int{
      {0, -1, 1}, // rock
      {1, 0, -1}, // paper
      {-1, 1, 0}, // scissors
    }
    scoreArr = []int {
      rockScore,
      paperScore,
      scissorScore,
    }
    resultMap = map[int]int {
      0: drawScore,
      -1: lostScore,
      1: winScore,
    }
  )
  m1Move := idMove(m1)
  m1Result := moveMatrix[idMove(m1)][idMove(m2)]
  // Game Result
  score = score + resultMap[m1Result]
  // Move Score
  score = score + scoreArr[m1Move]
  return m1Result, score
}

func main(){
  input := aoc22.LoadInputs("inputs.txt")
  for _, moves := range input {
    elf, player := loadMove(moves)
    _, score := winningMovePair(player, elf)
    totalScore = totalScore + score
  }
  fmt.Println(totalScore)
}