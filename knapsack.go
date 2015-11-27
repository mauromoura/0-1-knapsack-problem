package main

import "fmt"

func main() {
	values := []int{0, 2, 4, 2, 1, 10}
	weights := []int{0, 1, 12, 2, 1, 4}
	W := 10
	//values := []int{0, 30, 14, 16, 9}
	//weights := []int{0, 6, 3, 4, 2}
	//W := 10

	resultRecursive := knapsackRecursive(weights, values, len(weights)-1, W)
	resultDynamic := knapsackDynamic(weights, values, len(weights)-1, W)
	fmt.Println("recursive optimal value -> ", resultRecursive)
	fmt.Println("dynamic programming optimal value -> ", resultDynamic[len(weights)-1][W])
	fmt.Println("Optimal Solution through dynamic programming:")
	ra, rb := showOptimalSolution(values, weights, resultDynamic, W, len(weights)-1)
	fmt.Println(ra, rb)
}

func knapsackRecursive(weights []int, values []int, n int, W int) int {
	if n == 0 || W == 0 {
		return 0
	}
	without := knapsackRecursive(weights, values, n-1, W)
	if weights[n] > W {
		return without
	}
	withim := values[n] + knapsackRecursive(weights, values, n-1, W-weights[n])
	if withim > without {
		return withim
	}
	return without
}

func knapsackDynamic(weights []int, values []int, n int, W int) [][]int {
	m := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		m[i] = make([]int, W+1)
	}
	for i := 0; i <= W; i++ {
		m[0][i] = 0
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= W; j++ {
			if weights[i] > j {
				m[i][j] = m[i-1][j]
			} else if m[i-1][j] > m[i-1][j-weights[i]]+values[i] {
				m[i][j] = m[i-1][j]
			} else {
				m[i][j] = m[i-1][j-weights[i]] + values[i]
			}
		}
	}
	return m
}

func showOptimalSolution(values []int, weights []int, m [][]int, W int, n int) (int, []int) {
	finalValue := 0
	result := []int{}
	for W > 0 && n > 0 {
		if m[n][W] != m[n-1][W] {
			result = append(result, n)
			W = W - weights[n]
			finalValue += values[n]
			n--
		} else {
			n--
		}
	}
	return finalValue, result
}
