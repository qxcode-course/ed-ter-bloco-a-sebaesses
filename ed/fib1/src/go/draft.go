package main

import "fmt"

func totalCoelhos(n, k int) int {
    if n == 1 || n == 2 {
        return 1
    }
    f1, f2 := 1, 1
    var atual int

    for i := 3; i <= n; i++ {
        atual = f2 + k*f1
        f1 = f2
        f2 = atual
    }

    return f2
}

func main() {
    var n, k int
    if _, err := fmt.Scan(&n, &k); err != nil {
        return
    }

    fmt.Println(totalCoelhos(n, k))
}