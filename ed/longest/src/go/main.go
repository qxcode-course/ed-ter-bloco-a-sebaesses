package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])

	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}

	var dfs func(r, c, m, n int) int
	dfs = func(r, c, m, n int) int {
		if r < 0 || r >= m || c < 0 || c >= n {
			return 0
		}

		if memo[r][c] != 0 {
			return memo[r][c]
		}

		maxPath := 1
		dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]

			if nr >= 0 && nr < m && nc >= 0 && nc < n {
				if matrix[nr][nc] > matrix[r][c] {
					path := 1 + dfs(nr, nc, m, n)
					if path > maxPath {
						maxPath = path
					}
				}
			}
		}

		memo[r][c] = maxPath
		return maxPath
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			path := dfs(i, j, m, n)
			if path > ans {
				ans = path
			}
		}
	}

	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}