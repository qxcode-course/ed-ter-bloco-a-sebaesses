package main

import "fmt"

func calcularDinheiro(n int, c float64) float64 {
    dinheiro := 0.0

    for i := 0; i < n; i++ {
        dinheiro = (dinheiro + c) / 2
    }
    
    return dinheiro
}

func main() {
    var n int
    var c float64
    if _, err := fmt.Scan(&n, &c); err != nil {
        return
    }

    fmt.Printf("%.2f\n", calcularDinheiro(n, c))
}