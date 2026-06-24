package main

import "fmt"

var memo = make(map[int]int)

func ativos(n, k int) int {
    if n <= k {
        return 1
    }

    if val, ok := memo[n]; ok {
        return val
    }

    parte1 := (n+1)/2
    parte2 := n/2

    resultado := ativos(parte1, k) + ativos(parte2, k)

    memo[n] = resultado
    return resultado
}

func main() {
    var n, k int
    if _, err := fmt.Scan(&n, &k); err != nil {
        return
    }

    fmt.Println(ativos(n,k))

}