package main

import (
	"fmt"
	"math"
)

func getRoots(a, b, c float32) (isRoots bool, x1, x2 float32) {
	isRoots = true
	d := b*b - 4*a*c // Дискриминант
	if d > 0 {
		y := float32(math.Sqrt(float64(d)))
		x1 = (-1*b + y) / (2 * a)
		x2 = (-1*b - y) / (2 * a)
	} else if d == 0 {
		x1 = (b / 2 * a) * -1
		x2 = x1
	} else {
		isRoots = false
	}
	return isRoots, x1, x2
}

func main() {
	var a, b, c, x1, x2 float32

	fmt.Println("ax^2 + bx + c = 0")
	fmt.Print("Enter a = ")
	fmt.Scan(&a)
	fmt.Print("Enter b = ")
	fmt.Scan(&b)
	fmt.Print("Enter c = ")
	fmt.Scan(&c)
	isRoots, x1, x2 := getRoots(a, b, c)
	if !isRoots {
		_, _ = x1, x2
		fmt.Println("No roots")
		return
	}
	fmt.Printf("Founded roots x1 = %v, x2 = %v", x1, x2)
}
