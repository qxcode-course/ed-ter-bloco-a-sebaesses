package main

import "fmt"

func pares(n int) int {
    if n == 1 || n == 2 {
        return 1
    }

    if n == 3 || n == 4 {
        return 2
    }

    return pares(n-1) + pares(n-2) - pares(n-4)
}

func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil {
        return
    }

    fmt.Println(pares(n))
}