package main
import "fmt"

func main() {
	
	var n int
	fmt.Scan(&n)

	fila := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	var m int
	fmt.Scan(&m)

	sairam := make(map[int]bool)
	for i := 0; i < m; i++ {
		var desistente int
		fmt.Scan(&desistente)
		sairam[desistente] = true
	}

	for i := 0; i < n; i++ {
		if sairam[fila[i]] == false {
			fmt.Printf("%d ", fila[i])
		}
	}
	fmt.Println()
}

