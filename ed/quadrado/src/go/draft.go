package main

import "fmt"

func calcularQuadrado(n int) int {
    if n == 1 {
        fmt.Println("1^2 = 1")
        return 1
    }
    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?\n", n, n-1, n-1)
    prev := calcularQuadrado(n-1)
    curr := prev + 2*(n-1) + 1
    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = %d\n", n, n-1, n-1, curr)
    return curr
}

func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil {
        return
    }
    calcularQuadrado(n)
}