package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	for _, capacity := range []uint64{
		1,
		2,
		3,
		30,
		100,
		1000,
		10000,
	} {
		iterations, estimate := monteCarlo(100000000, capacity, 0.0001)

		fmt.Printf(
			"capacity: %-6d| iterations: %-9d| estimate: %f\n",
			capacity,
			iterations,
			estimate,
		)
	}
}

// monteCarlo runs at most iterationsMax experiments and retuns the average estimate.
// The function returns early if estimates convege within epsilon.
func monteCarlo(iterationsMax uint64, capacity uint64, epsilon float64) (uint64, float64) {
	var (
		total uint64

		// circular array of 10 latest estimates (computed every 1000 iterations)
		latestEstimates   [10]float64
		latestEstimateIdx int
	)

	withinEpsilon := func(latestEstimates []float64, latestEstimateIdx int, epsilon float64) bool {
		for i := 0; i < len(latestEstimates); i++ {
			if math.Abs(latestEstimates[i]-latestEstimates[latestEstimateIdx]) > epsilon {
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

			if withinEpsilon(latestEstimates[:], latestEstimateIdx, epsilon) {
				fmt.Println(latestEstimates)
				return i, latest
			}
		}
	}

	return iterationsMax, float64(total) / float64(iterationsMax)
}

// calculateOne runs a single calculation for the given number of oranges in each box (capacity)
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
