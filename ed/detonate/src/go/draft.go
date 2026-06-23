package main

import "fmt"

func maxDetonation(bombs [][]int) int {
    n := len(bombs)
    graph := make([][]int, n)
    
    for i:= 0; i < n; i++ {
        for j:= 0; j < n; j++ {
            if i == j {
                continue
            }
            dx := int64(bombs[i][0] - bombs[j][0])
            dy := int64(bombs[i][1] - bombs[j][1])
            distSq := dx*dx + dy*dy
            rSq := int64(bombs[i][2]) * int64(bombs[i][2])
            if distSq <= rSq {
                graph[i] = append(graph[i], j)
            }
        }
    }
    
    maxDetonated := 0
        for i := 0; i < n; i++ {
            visited := make([]bool, n)
            count := 0
            
            var dfs func(node int)
            dfs = func(node int) {
                visited[node] = true
                count++

                for _, vizinho := range graph[node] {
                    if !visited[vizinho] {
                        dfs(vizinho)
                    }
                }
            }

            dfs(i)
            if count > maxDetonated {
                maxDetonated = count
            }
        }

        return maxDetonated
}

func main() {
    var n, cols int
    if _, err := fmt.Scan(&n, &cols); err != nil {
        return
    }

    bombs := make([][]int, n)
    for i := 0; i < n; i++ {
        bombs[i] = make([]int, 3)
        fmt.Scan(&bombs[i][0], &bombs[i][1], &bombs[i][2])
    }

    fmt.Println(maxDetonation(bombs))
}