package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(grid [][]byte, r, c int) {
	grid[r][c] = '0'

	if c+1 < len(grid[0]) && grid[r][c+1] == '1' {
		dfs(grid, r, c+1)
	}
	if r+1 < len(grid) && grid[r+1][c] == '1' {
		dfs(grid, r+1, c)
	}
}

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
