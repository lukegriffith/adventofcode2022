package main

import "testing"

func TestIDMove(t *testing.T) {
  if idMove(move("X")) != 0 {
    t.Fail()
  }
  if idMove(move("A")) != 0 {
    t.Fail()
  }
  if idMove(move("B")) != 1 {
    t.Fail()
  }
  if idMove(move("Y")) != 1 {
    t.Fail()
  }
  if idMove(move("C")) != 2 {
    t.Fail()
  }
  if idMove(move("Z")) != 2 {
    t.Fail()
  }
}

func TestMoveLogic(t *testing.T) {
  result, score := winningMovePair(move("Y"), move("A"))
  t.Log(result,score)
  if result != 1 {
    t.Fail()
  }
  if score != 8 {
    t.Fail()
  }

  result, score = winningMovePair(move("X"), move("B"))
  t.Log(result,score)
  if result != -1 {
    t.Fail()
  }
  if score != 1 {
    t.Fail()
  }

  result, score = winningMovePair(move("C"), move("Z"))
  t.Log(result,score)
  if result != 0 {
    t.Fail()
  }
  if score != 6 {
    t.Fail()
  }
}