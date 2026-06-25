package main

import "fmt"

func colocar(matriz [][]rune, l, c int, val rune) bool {
    n := len(matriz)
    for i := 0; i < n; i++ {
        if matriz[l][i] == val || matriz[i][c] == val {
            return false
        }
    } 

    tamanhoBloco := 2
    if n == 9 {
        tamanhoBloco = 3
    }

    inicioL := (l/tamanhoBloco) * tamanhoBloco
    inicioC := (c/tamanhoBloco) * tamanhoBloco

    for i := 0; i < tamanhoBloco; i++ {
        for j := 0; j < tamanhoBloco; j++ {
            if matriz[inicioL+i][inicioC+j] == val {
				return false
			}
        }
    }

    return true
}

func resolver(matriz [][]rune, index int) bool {
    n := len(matriz)
    if index == n*n {
        return true
    }

    l := index / n
    c := index % n

    if matriz[l][c] != '.' {
        return resolver(matriz, index+1)
    }

    for v := '1'; v <= rune('0'+n); v++ {
        if colocar(matriz, l, c, v) {
            matriz[l][c] = v
            if resolver(matriz, index+1) {
                return true
            }

            matriz[l][c] = '.'
        }
    }

    return false
}

func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil {
        return
    }
    matriz := make([][]rune, n)
    for i := 0; i < n; i++ {
        var linha string
        fmt.Scan(&linha)
        matriz[i] = []rune(linha)
    }

    resolver(matriz, 0)
    for i := 0; i < n; i++ {
        fmt.Println(string(matriz[i]))
    }
}