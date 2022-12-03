package main

import "testing"

func TestSplitBackpack (t *testing.T) {
  s1, s2 := splitBackpack("aaabbb")
  t.Log(s1, "split", s2)
  if s1 != "aaa"{
    t.Fail()
  }
  if s2 != "bbb" {
    t.Fail()
  }

  s1, s2 = splitBackpack("aaaabbb")
  t.Log(s1, "split", s2)
  if s1 != "aaa" {
    t.Fail()
  }
  if s2 != "abbb" {
    t.Fail()
  }
}

func TestStringRange(t *testing.T) {
  for _, i := range "hello" {
    if i != 'h' {
      t.Fail()
    }
    return
  }
}

func TestGetItemValue(t *testing.T) {
  i := item('a')
  t.Log(i.getValue())
  if i.getValue() != 1 {
    t.Fail()
  }
  i = item('z')
  t.Log(i.getValue())
  if i.getValue() != 26 {
    t.Fail()
  }
  i = item('A')
  t.Log(i.getValue())
  if i.getValue() != 27 {
    t.Fail()
  }
  i = item('Z')
  t.Log(i.getValue())
  if i.getValue() != 52 {
    t.Fail()
  }
}

