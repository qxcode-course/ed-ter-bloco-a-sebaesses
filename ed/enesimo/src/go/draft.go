package main

import "fmt"

func ehPrimo(num int, div int) bool {
    if num <= 1 {
        return false
    }
    if div == 1 {
        return true
    }
    if num%div == 0 {
        return false
    }
    return ehPrimo(num, div-1)
}

func buscaPrimo(atual int, enc int, n int) int {
    if ehPrimo(atual, atual-1) {
        enc++
    }
    if enc == n {
        return atual
    }

    return buscaPrimo(atual+1, enc, n)

}

func enesimoPrimo(n int) int {
    return buscaPrimo(2, 0, n)
}

func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil {
        return
    }

    fmt.Println(enesimoPrimo(n))  
}