package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l int
	c int
}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	stack.Push(Pos{l, c})

	for !stack.IsEmpty() {
		curr := stack.Pop()
		r := curr.l
		col := curr.c

		if r >= 0 && r < len(grid) && col >= 0 && col < len(grid[0]) && grid[r][col] == '#' {
			grid[r][col] = 'o'
			
			stack.Push(Pos{r - 1, col})
			stack.Push(Pos{r + 1, col})
			stack.Push(Pos{r, col - 1})
			stack.Push(Pos{r, col + 1})
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
