package main

import "fmt"

func imprimir(n string) {
    if len(n) == 0 {
        return
    }

    imprimir(n[1:])
    fmt.Println(n)
}

func main() {
    var n string
	
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	imprimir(n)
}