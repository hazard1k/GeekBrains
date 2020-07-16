package main

import (
	"fmt"
	"math"
)

var infoString string = `
    C
    .
    |\
    | \
   a|  \ c=?
    |   \
    |____\
    A  b  C
`

var (
	a, b float64
)

func main() {
	fmt.Println(infoString)
	fmt.Print("Enter a = ")
	fmt.Scan(&a)
	fmt.Print("Enter b = ")
	fmt.Scan(&b)
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	fmt.Printf("c = %v", c)
}
