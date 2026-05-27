package main

import "fmt"

func isValid(arr []byte, idx int, L int, char byte) bool {
	inicio := idx - L
	if inicio < 0 {
		inicio = 0
	}
	for i := inicio; i < idx; i++ {
		if arr[i] == char {
			return false
		}
	}
	fim := idx + L
	if fim >= len(arr) {
		fim = len(arr) - 1
	}
	for i := idx + 1; i <= fim; i++ {
		if arr[i] == char {
			return false
		}
	}
	return true
}

func main() {
	var seq string
	var L int

	for {
		n, _ := fmt.Scan(&seq, &L)
		if n != 2 {
			break
		}
		arr := []byte(seq)
		var solve func(idx int) bool
		solve = func(idx int) bool {
			if idx == len(arr) {
				return true
			}
			if arr[idx] != '.' {
				return solve(idx + 1)
			}

			for d := 0; d <= L; d++ {
				if isValid(arr, idx, L, '0'+byte(d)) {
					arr[idx] = '0' + byte(d)
					if solve(idx + 1) {
						return true
					}
					arr[idx] = '.'
				}
			}
			return false
		}
		solve(0)
		fmt.Println(string(arr))
	}
}