package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func insertRecursive(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}

	if value < node.Value {
		node.Left = insertRecursive(node.Left, value)
	} else if value > node.Value {
		node.Right = insertRecursive(node.Right, value)
	}
	return node
}

func BstInsert(values []int) *Node {
	var root *Node

	for _, value := range values {
		root = insertRecursive(root, value)
	}
	return root
}


func serializeAux(node *Node, serial *[]string) {
	if node == nil {
		*serial = append(*serial, "#")
		return
	}
	*serial = append(*serial, strconv.Itoa(node.Value))
	serializeAux(node.Left, serial)
	serializeAux(node.Right, serial)
}

// Dica: crie um vetor compartilhado e vá preenchendo conforme anda na recursão
// Depois use o strings.Join para gerar o serial
func Serialize(root *Node) string {
	var result []string
	serializeAux(root, &result)
	return strings.Join(result, " ")
}

// -----------------------------------------------------------------------------------
func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}
	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, history+"r")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	values := make([]int, 0, len(parts))
	for _, elem := range parts {
		v, err := strconv.Atoi(elem)
		if err == nil {
			values = append(values, v)
		}
	}
	root := BstInsert(values)
	BShow(root, "") // Chama a função de impressão formatada
	fmt.Println(Serialize((root)))
}
