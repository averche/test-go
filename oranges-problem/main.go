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