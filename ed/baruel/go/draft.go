package main
import "fmt"

func main() {

    var totalAlbum int
	var totalCompradas int
	fmt.Scan(&totalAlbum)
	fmt.Scan(&totalCompradas)

	var album [51]int

	figurinhaRepetida := false
	figurinhaFalta := false

	for i := 0; i < totalCompradas; i++ {
		var figurinha int
		fmt.Scan(&figurinha)

		if album[figurinha] > 0 {
			if figurinhaRepetida == true {
				fmt.Print(" ")
			}
			fmt.Print(figurinha)
			figurinhaRepetida = true
		}
		album[figurinha]++
	}
	if figurinhaRepetida == false {
			fmt.Print("N")
	}
	fmt.Println()

	for i := 1; i <= totalAlbum; i++ {
		if album[i] == 0 {
			if figurinhaFalta == true {
				fmt.Print(" ")
			}
			fmt.Print(i)
			figurinhaFalta = true
		}
	}

	if figurinhaFalta == false { 
		fmt.Print("N")
	}
	fmt.Println()

}
