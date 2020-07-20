package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter the number: ")
	_, err := fmt.Scanf("%d", &number)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	res := "odd"
	if number%2 == 0 {
		res = "even"
	}

	fmt.Printf("The number %d is %v", number, res)
}
