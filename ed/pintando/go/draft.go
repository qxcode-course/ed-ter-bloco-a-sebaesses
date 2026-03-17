package main

import "fmt"
import "math"

func main(){

var a float64
var b float64
var c float64


fmt.Scan(&a)
fmt.Scan(&b)
fmt.Scan(&c)

p := (a + b + c) / 2.0

area := math.Sqrt(p *(p-a)*(p-b)*(p-c))

fmt.Printf("%.2f\n", area)

}
