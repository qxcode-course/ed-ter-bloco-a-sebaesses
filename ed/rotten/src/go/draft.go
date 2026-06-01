package main
import "fmt"

type Orange struct {
    r, c, min int
}

func orangesRotting(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    q := []Orange{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 2 {
                q = append(q, Orange{i, j, 0})
            }
        }
    }
    ans := 0
    for len(q) > 0 {
        o := q[0]
        q = q[1:]
        ans = o.min
        for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
            r, c := o.r+d[0], o.c+d[1]
            if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == 1 {
                grid[r][c] = 2
                q = append(q, Orange{r, c, o.min + 1})
            }
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                return -1
            }
        }
    }
    return ans
}  

func main() {
    var m, n int
    fmt.Scan(&m, &n)
    grid := make([][]int, m)
    for i := 0; i < m; i++ {
        grid[i] = make([]int, n)
        for j := 0; j < n; j++ {
            fmt.Scan(&grid[i][j])
        }
    }
    fmt.Println(orangesRotting(grid))
}