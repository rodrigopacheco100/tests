package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	grid := make([][]int, N)
	for i := 0; i < N; i++ {
		grid[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Scan(&grid[i][j])
		}
	}

	minAdditions, result := Solve(grid)

	fmt.Println(minAdditions)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(result[i][j])
		}
		fmt.Println()
	}
}

func Solve(grid [][]int) (int, [][]int) {
	N := len(grid)
	if N == 0 {
		return 0, nil
	}
	M := len(grid[0])

	cost1 := 0
	cost2 := 0

	result1 := make([][]int, N)
	result2 := make([][]int, N)

	for i := range N {
		result1[i] = make([]int, M)
		result2[i] = make([]int, M)

		for j := range M {
			if (i+j)%2 == 0 {
				if grid[i][j]%2 == 0 {
					result1[i][j] = grid[i][j]
				} else {
					result1[i][j] = grid[i][j] + 1
					cost1++
				}
				if grid[i][j]%2 == 1 {
					result2[i][j] = grid[i][j]
				} else {
					result2[i][j] = grid[i][j] + 1
					cost2++
				}
			} else {
				if grid[i][j]%2 == 1 {
					result1[i][j] = grid[i][j]
				} else {
					result1[i][j] = grid[i][j] + 1
					cost1++
				}
				if grid[i][j]%2 == 0 {
					result2[i][j] = grid[i][j]
				} else {
					result2[i][j] = grid[i][j] + 1
					cost2++
				}
			}
		}
	}

	if cost1 <= cost2 {
		return cost1, result1
	}
	return cost2, result2
}
