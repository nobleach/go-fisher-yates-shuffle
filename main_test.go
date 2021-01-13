package main

import "testing"

func TestShuffle(t *testing.T) {
  q := []int{1, 2, 3, 4, 5}
  staticSeed := int64(4) // by roll of dice
  Shuffle(q, staticSeed)

  if q[0] != 1 {
    t.Errorf("Expected 1 but got %v", q[0])
  }

  if q[4] != 3 {
    t.Errorf("Expected 3 but got %v", q[4])
  }
}
