package main

import "fmt"

func calcularPecas(n int) int {
    return n*(n+2)
}

func main() {
    var n int
	
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

    fmt.Println(calcularPecas(n))
}