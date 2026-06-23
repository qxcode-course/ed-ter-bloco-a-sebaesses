package main

import "fmt"
import "strconv"
import "strings"

func ehPrimo(num int) bool {
    if num <= 1 {
        return false
    }
    for i:= 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }

    return true
}

func geraPrimos(n int) []int {
    var primos []int
    atual := 2

    for len(primos) < n {
        if ehPrimo(atual) {
            primos = append(primos, atual)
        }
        atual++
    }

    return primos
    }

func formatArray(arr []int) string {
    strs := make([]string, len(arr))
    for i, v := range arr {
        strs[i] = strconv.Itoa(v)
    }
    return "[" + strings.Join(strs, ", ") + "]"
}

func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil {
        return
    }

    primos := geraPrimos(n)
    fmt.Println(formatArray(primos))
}