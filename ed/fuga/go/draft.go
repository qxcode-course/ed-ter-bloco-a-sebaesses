package main
import "fmt"

func main() {


    var H, P, F, D int

fmt.Scan(&H, &P, &F, &D)

for {
    if F == H {
			fmt.Println("S")
			break
		}
		if F == P {
			fmt.Println("N")
			break 
		}

		F = F + D

		if F > 15 {
			F = 0
		} else if F < 0 {
			F = 15
		}
	}
}


