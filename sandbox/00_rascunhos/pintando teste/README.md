package main

import "fmt"
import "math"

func main(){

var a int 3.0
var b int 4.0
var c int 5.0

p := (a + b + c) / 2.0

area := math.Sqrt(p *(p-a)*(p-b)*(p-c))

fmt.printf("%.2f\n", area)

}
