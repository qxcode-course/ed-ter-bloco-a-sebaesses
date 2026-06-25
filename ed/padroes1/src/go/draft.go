package main

import "fmt"

func calcularPecas(n int) int {
    return 8*n + 12
}

func main() {
     var n int
	
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

    fmt.Println(calcularPecas(n))
}