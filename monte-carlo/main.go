package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	for _, capacity := range []uint64{
		30,
		100,
		1000,
		10000,
	} {
		iterations, estimate := monteCarlo(100000000, capacity, 0.001)

		fmt.Printf(
			"capacity: %-9d, operations: %-9d, estimate: %f\n",
			capacity,
			iterations,
			estimate,
		)
	}
}

func monteCarlo(iterationsMax uint64, capacity uint64, epsilon float64) (uint64, float64) {
	var (
		total             uint64
		latestEstimates   [10]float64
		latestEstimateIdx int
	)

	withinEpsilon := func(estimates []float64, epsilon float64) bool {
		for i := 1; i < len(estimates); i++ {
			if math.Abs(estimates[i]-estimates[i-1]) > epsilon {
				return false
			}
		}
		return true
	}

	for i := uint64(0); i < iterationsMax; i++ {
		total += calculateOne(
			time.Now().UnixNano()+int64(i),
			capacity,
		)

		if i%1000 == 0 {
			latest := float64(total) / float64(i+1)
			latestEstimates[latestEstimateIdx] = latest
			latestEstimateIdx = (latestEstimateIdx + 1) % 10

			// if the last 3 estimates are within epsilon, return early
			if withinEpsilon(latestEstimates[:], epsilon) {
				return i, latest
			}
		}
	}

	return iterationsMax, float64(total) / float64(iterationsMax)
}

func calculateOne(seed int64, capacity uint64) uint64 {
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
