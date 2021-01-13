package main

import (
  "fmt"
	"math/rand"
	"time"
)

func main() {
  q := []int{1, 2, 3, 4, 5}
  fmt.Println(q);
  Shuffle(q, time.Now().UnixNano())
  fmt.Println(q);
}

func Shuffle(l []int, seed int64) {
  source := rand.NewSource(seed)
  r := rand.New(source)

  for i := range l {
    newPosition := r.Intn(len(l) - 1)

    l[i], l[newPosition] = l[newPosition], l[i]
  }
}
