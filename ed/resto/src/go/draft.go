package main
import "fmt"

func divisao(num int)  {
    if num == 0 {
        return
    }
    
    quociente := num / 2
    resto := num % 2

    divisao(quociente)
    
    fmt.Printf("%d %d\n", quociente, resto)
}

func main() {
    var num int
    fmt.Scan(&num)
    
    if num > 0 {
        divisao(num)
    }
}
