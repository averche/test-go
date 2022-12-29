package main

import "fmt"

func main() {
	for _, n := range []int{
		1,
		2,
		3,
		30,
		100,
		1000,
		10000,
	} {
		fmt.Printf(
			"capacity: %-6d| expected oranges: %f\n",
			n,
		 	expectedOranges(n),
		)
	}
}

// dynamic programming
func expectedOranges(n int) float64 {
	// Initialize the dynamic programming table
	table := make([][]float64, n+1)
	for i := range table {
		table[i] = make([]float64, n+1)
	}

	// Base case: If one of the boxes has 0 oranges, the expected value is the number of oranges in the other box
	for i := 0; i <= n; i++ {
		table[i][0] = float64(i)
		table[0][i] = float64(i)
	}

	// Solve the subproblems using dynamic programming
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			// The expected value is the average of the expected values when an orange is taken from either box
			table[i][j] = 0.5 * (table[i-1][j] + table[i][j-1])
		}
	}

	// Return the expected value for the given number of oranges in each box
	return table[n][n]
}

// recursive + memo
func expectedOrangesRecursive(n int) float64 {
	var sum float64
	var count float64
	memo := make(map[string]float64)
	buildTree(n, n, 1, &sum, &count, memo)
	return sum / count
}

func buildTree(a int, b int, prob float64, sum *float64, count *float64, memo map[string]float64) {
	key := fmt.Sprintf("%d,%d", a, b)
	if val, ok := memo[key]; ok {
		*sum += val * prob
		*count += prob
		return
	}
	if a == 0 {
		*sum += prob * float64(b)
		*count += prob
		memo[key] = float64(b)
		return
	}
	if b == 0 {
		*sum += prob * float64(a)
		*count += prob
		memo[key] = float64(a)
		return
	}
	buildTree(a-1, b, prob/2, sum, count, memo)
	buildTree(a, b-1, prob/2, sum, count, memo)
}

// expectedOrangesMonteCarlo runs at most iterationsMax experiments and retuns the average estimate.
// The function returns early if estimates convege within epsilon.
func expectedOrangesMonteCarlo(iterationsMax uint64, capacity uint64, epsilon float64) (uint64, float64) {
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
