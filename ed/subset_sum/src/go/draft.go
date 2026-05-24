package main
import "fmt"

func subsetSum(arr []int, k int, index int) bool {
        if k == 0 {
            return true
        }
        if index >= len(arr) {
            return false
        }
        // Include current element
        if subsetSum(arr, k-arr[index], index+1) {
            return true
        }
        // Exclude current element
        return subsetSum(arr, k, index+1)
    }

func main() {
    var n, k int
    
    if _, err := fmt.Scan(&n, &k); err != nil {
        return
}
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    if subsetSum(arr, k, 0) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}