package main

import "fmt"

func subirEscada(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 1
    }

    if n == 3 {
        return 2
    }

    return subirEscada(n-1) + subirEscada(n-3)
}

func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil {
        return
    }

    fmt.Println(subirEscada(n))
}