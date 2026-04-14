package main
import "fmt"

func main() {

    var n int
    fmt.Scan(
        &n)
    
    semPar := make(map[int]int)

    pares := 0

    for i := 0; i < n; i++ {
        var animal int
        fmt.Scan(&animal)

        parIdeal := -animal
        
        if semPar [parIdeal] > 0 {
            pares++
            semPar[parIdeal]--
        } else {
            semPar[animal]++
        }
    }

    fmt.Println(pares)

}
