package main

import "fmt"

func imprimirRoda(roda []int, indiceEspada int) {
	fmt.Print("[ ")
	for i := 0; i < len(roda); i++ {
		if i == indiceEspada {
			fmt.Printf("%d> ", roda[i])
		} else {
			fmt.Printf("%d ", roda[i])
		}
	}
	fmt.Println("]")
}

func main() {
	var n, e int
	fmt.Scan(&n, &e)

	roda := make([]int, n)
	for i := 0; i < n; i++ {
		roda[i] = i + 1
	}

	indiceEspada := e - 1

	for len(roda) > 1 {
		imprimirRoda(roda, indiceEspada)

		indiceMorto := (indiceEspada + 1) % len(roda)

		roda = append(roda[:indiceMorto], roda[indiceMorto+1:]...)

		indiceEspada = indiceMorto % len(roda)
	}

	imprimirRoda(roda, indiceEspada)
}
