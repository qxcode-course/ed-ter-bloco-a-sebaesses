package main

import "fmt"

func binominal(n, k int) int {
    if k == 0 || k == n {
        return 1
    }
    
    return binominal(n-1, k-1) + binominal(n-1, k)
}

func main() {
    var n, k int
    if _, err := fmt.Scan(&n, &k); err != nil {
        return
    }

    fmt.Println(binominal(n, k))
    
}