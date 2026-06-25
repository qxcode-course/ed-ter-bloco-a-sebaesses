package main

import "fmt"

func calcularPontos(n, m int) int {
    if m == 1 {
        return 1
    }

    return calcularPontos(n, m-1) + (n-2)*(m-1) + 1
}

func main() {
    var n, m int
    if _, err := fmt.Scan(&n, &m); err != nil {
        return
    }

    fmt.Println(calcularPontos(n, m ))
}