package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func occurr(vet []int) []Pair {
	counts := make(map[int]int)
	for _, v := range vet {
		counts[abs(v)]++
	}

	var keys []int
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	
	var result []Pair
	for _, k := range keys {
		result = append(result, Pair{One: k, Two: counts[k]})
	}
	return result
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return nil
	}

	var result []Pair
	currentStress := abs(vet[0])
	count := 1

	for i := 1; i < len(vet); i++ {
		s := abs(vet[i])
		if s == currentStress {
			count++
		} else {
			result = append(result, Pair{One: currentStress, Two: count})
			currentStress = s
			count = 1
		}
	}
	result = append(result, Pair{One: currentStress, Two: count})

	return result
}

func mnext(vet []int) []int {
	if len(vet) == 0 {
		return nil
	}
	result := make([]int, len(vet))

	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			nextToWoman := false
			if i > 0 && vet[i-1] < 0 {
				nextToWoman = true
			}
			if i < len(vet)-1 && vet[i+1] < 0 {
				nextToWoman = true
			}
			if nextToWoman {
				result[i] = 1
			}
		}
	}
	return result
}

func alone(vet []int) []int {
	if len(vet) == 0 {
		return nil
	}
	result := make([]int, len(vet))

	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			nextToWoman := false
			if i > 0 && vet[i-1] < 0 {
				nextToWoman = true
			}
			if i < len(vet)-1 && vet[i+1] < 0 {
				nextToWoman = true
			}
			if !nextToWoman {
				result[i] = 1
			}
		}
	}
	return result
}


func couple(vet []int) int {
	casados := make([]bool, len(vet)) 
	casais := 0

	for i := 0; i < len(vet); i++ {
		if casados[i] {
			continue 
		}
		
		for j := i + 1; j < len(vet); j++ {
			if !casados[j] && abs(vet[i]) == abs(vet[j]) && vet[i] != vet[j] {
				casados[i] = true
				casados[j] = true
				casais++
				break 
			}
		}
	}
	return casais
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	if pos+len(seq) > len(vet) {
		return false
	}
	
	for i := 0; i < len(seq); i++ {
		if vet[pos+i] != seq[i] {
			return false
		}
	}
	
	return true
}

func subseq(vet []int, seq []int) int {
	if len(seq) == 0 {
		return 0
	}
	
	for i := 0; i <= len(vet)-len(seq); i++ {
		if hasSubseq(vet, seq, i) {
			return i 
		}
	}
	
	return -1 
}

func erase(vet []int, posList []int) []int {
	toRemove := make(map[int]bool)
	for _, pos := range posList {
		toRemove[pos] = true
	}

	var result []int
	for i, v := range vet {
		if !toRemove[i] {
			result = append(result, v)
		}
	}
	
	if result == nil {
		return []int{} 
	}
	return result
}

func clear(vet []int, value int) []int {
	var result []int
	for _, v := range vet {
		if v != value {
			result = append(result, v)
		}
	}
	
	if result == nil {
		return []int{} 
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
