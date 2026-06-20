package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	r, c int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	if !scanner.Scan() {
		return
	}
	dims := strings.Fields(scanner.Text())
	if len(dims) < 2 {
		return
	}
	rows, _ := strconv.Atoi(dims[0])

	maze := make([][]rune, rows)
	var start Pos

	for i := 0; i < rows; i++ {
		scanner.Scan()
		line := scanner.Text()
		maze[i] = []rune(line)
		for j, ch := range maze[i] {
			if ch == 'I' {
				start = Pos{r: i, c: j}
			}
		}
	}

	caminho := NewStack[Pos]()
	becos := NewStack[Pos]()
	caminho.Push(start)

	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, len(maze[i]))
	}

	dr := []int{-1, 0, 1, 0}
	dc := []int{0, 1, 0, -1}

	for !caminho.IsEmpty() {
		atual := caminho.Top()
		visited[atual.r][atual.c] = true

		if maze[atual.r][atual.c] == 'F' {
			break
		}

		var validos []Pos
		for i := 0; i < 4; i++ {
			nr, nc := atual.r+dr[i], atual.c+dc[i]
			if nr >= 0 && nr < rows && nc >= 0 && nc < len(maze[nr]) {
				if maze[nr][nc] != '#' && !visited[nr][nc] {
					validos = append(validos, Pos{r: nr, c: nc})
				}
			}
		}

		if len(validos) > 0 {
			caminho.Push(validos[0])
		} else {
			beco := caminho.Pop()
			becos.Push(beco)
		}
	}

	for !caminho.IsEmpty() {
		p := caminho.Pop()
		maze[p.r][p.c] = '.'
	}

	for i := 0; i < rows; i++ {
		fmt.Println(string(maze[i]))
	}
}