package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

func getMen(vet []int) []int {
	var homens []int
	for _, v := range vet {
		if v > 0 {
			homens = append(homens, v)
		}
	}
	return homens
}

func getCalmWomen(vet []int) []int {
	var calmas []int
	for _, v := range vet {
		if v < 0 && v > -10 {
			calmas = append(calmas, v)
		}
	}
	return calmas
}

func sortVet(vet []int) []int {
	sort.Ints(vet)
	return vet
}

func sortStress(vet []int) []int {
	sort.Slice(vet, func(i, j int) bool {
		return abs(vet[i]) < abs(vet[j])
	})
	return vet
}

func reverse(vet []int) []int {
	var invertido []int
	for i := len(vet) - 1; i >= 0; i-- {
		invertido = append(invertido, vet[i])
	}
	return invertido
}

func unique(vet []int) []int {
	var unicos []int
	vistos := make(map[int]bool) 

	for _, v := range vet {
		if vistos[v] == false {
			unicos = append(unicos, v)
			vistos[v] = true 
		}
	}
	return unicos
}

func repeated(vet []int) []int {
	var repetidos []int
	vistos := make(map[int]bool)

	for _, v := range vet {
		if vistos[v] == true {
			repetidos = append(repetidos, v)
		} else {
			vistos[v] = true 
		}
	}
	return repetidos
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

