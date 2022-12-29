package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(calculateOne(time.Now().UnixNano(), 30))
}

func calculateOne(seed int64, capacity uint32) uint32 {
	// intialice both to the given capacity
	oranges1, oranges2 := capacity, capacity

	rng := rand.New(rand.NewSource(seed))

	for {
		if rng.Uint64()%2 == 0 {
			oranges1--
		} else {
			oranges2--
		}

		if oranges1 == 0 {
			return oranges2
		}
		if oranges2 == 0 {
			return oranges1
		}
	}
}
