package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(board [][]byte, r, c int) {
	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) || board[r][c] != 'O' {
		return
	}

	board[r][c] = 'S'

	dfs(board, r-1, c)
	dfs(board, r+1, c)
	dfs(board, r, c-1)
	dfs(board, r, c+1)
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	rows := len(board)
	cols := len(board[0])

	for j := 0; j < cols; j++ {
		if board[0][j] == 'O' {
			dfs(board, 0, j)
		}
		if board[rows-1][j] == 'O' {
			dfs(board, rows-1, j)
		}
	}

	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][cols-1] == 'O' {
			dfs(board, i, cols-1)
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'S' {
				board[i][j] = 'O'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
