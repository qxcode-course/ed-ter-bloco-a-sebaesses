package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	var dfs func(r, c, index int) bool
	dfs = func(r, c, index int) bool {
		if index == len(word) {
			return true
		}
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != word[index] {
			return false
		}
		temp := grid[r][c]
		grid[r][c] = '#'
		found := dfs(r+1, c, index+1) || dfs(r-1, c, index+1) || dfs(r, c+1, index+1) || dfs(r, c-1, index+1)
		grid[r][c] = temp
		return found
	}
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
