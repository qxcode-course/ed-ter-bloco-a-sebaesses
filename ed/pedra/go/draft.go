 package main
import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

    var n int
    fmt.Scan(&n)

    var melhorIndice int
    var menorDiferenca int

    teveGanhador := false

    for i := 0; i < n; i++ {
        var a, b int
        fmt.Scan(&a, &b)

        if a >= 10 && b >= 10 {
            diferenca := abs(a - b)

            if teveGanhador == false || diferenca < menorDiferenca{
                menorDiferenca = diferenca
                melhorIndice = i
                teveGanhador = true
            }
        }

    }   

    if teveGanhador == false {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(melhorIndice)
    }

}
