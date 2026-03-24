package main

import "fmt"

type Ponto struct {
	x int
	y int
}

func main() {
	var q int
	var d string
	fmt.Scan(&q, &d)

	cobra := make([]Ponto, q)

	for i := 0; i < q; i++ {
		fmt.Scan(&cobra[i].x, &cobra[i].y)
	}

	for i := q - 1; i > 0; i-- {
		cobra[i] = cobra[i-1] 
	}

	
	if d == "L" {
		cobra[0].x-- 
	} else if d == "R" {
		cobra[0].x++ 
	} else if d == "U" {
		cobra[0].y--
 	} else if d == "D" {
		cobra[0].y++ 
	}

	for i := 0; i < q; i++ {
		fmt.Println(cobra[i].x, cobra[i].y)
	}
}