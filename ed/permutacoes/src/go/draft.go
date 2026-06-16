package main
import (
    "fmt"
    "sort"
)

func permute(chars []rune, current []rune, visited []bool) {
    if len(current) == len(chars) {
        fmt.Println(string(current))
        return
    }

    for i := 0; i < len(chars); i++ {
        if !visited[i] {
            visited[i] = true
            current = append(current, chars[i])
            permute(chars, current, visited)
            current = current[:len(current)-1]
            visited[i] = false
        }
    }
}

func main() {
    var s string
    if _, err := fmt.Scan(&s); err != nil {
        return
    }

    chars := []rune(s)
    sort.Slice(chars, func(i, j int) bool {
        return chars[i] < chars[j]
    })

    visited := make([]bool, len(chars))
    var current []rune
    permute(chars, current, visited)
    
}