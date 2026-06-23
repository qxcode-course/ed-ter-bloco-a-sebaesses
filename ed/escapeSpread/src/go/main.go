package main

import (
	"fmt"
)

type pair struct{
	r, c int
}

func maximoMinutos(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	firetime := make ([][]int, m)
	var fireQ []pair

	for i := 0; i < m; i++ {
		fireTime[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fireTime[i][j] = 1e9 // Preenchemos com "Infinito"
			if grid[i][j] == 1 {
				fireTime[i][j] = 0 // Fogo começa no minuto 0
				fireQ = append(fireQ, pair{i, j})
			}
		}
	}

	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(fireQ) > 0 {
		curr := fireQ[0]
		fireQ := 
	}
}

func main() {
	var m, n int
	if _, err := fmt.Scan(&m, &n); err != nil {
		return
	}

	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&grid[i][i])
		}
	}

	fmt.Println(maximoMinutos(grid))
}

