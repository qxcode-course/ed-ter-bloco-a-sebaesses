package main

import (
	"fmt"
	"strings"
)

func diagonal(s string, k int) {
    if len(s) == 0 {
        return
    }

    espacos := strings.Repeat(" ", k)
	letra := string(s[0])
	
	fmt.Printf("%s%s\n", espacos, letra)
	
	diagonal(s[1:], k+1)
}

func main() {
    var n string
	
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	diagonal(n, 0)
}