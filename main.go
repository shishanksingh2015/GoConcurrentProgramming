package main

import (
	"fmt"
	"math/rand"
	"time"
)

// var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 13; i++ {
		id := rnd.Intn(13) + 1
		fmt.Printf("id = %v\n ", id)
	}
}
