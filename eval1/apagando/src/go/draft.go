package main
import "fmt"

func main() {
	
	var n int
	fmt.Scan(&n)

	idFila := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&idFila[i])
	}

	var m int
	fmt.Scan(&m)

	saiu := make(map[int]bool)
	for i := 0; i < m; i++ {
		var deixaramFila int
		fmt.Scan(&deixaramFila)
		saiu[deixaramFila] = true
	}

	for i := 0; i < n; i++ {
		if saiu[idFila[i]] == false {
			fmt.Printf("%d ", idFila[i])
		}
	}
	fmt.Println()
}

