package main

import "fmt"

func main() {
	var number int
	var res string = "No"
	fmt.Print("Enter the number: ")
	_, err := fmt.Scanf("%d", &number)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if number%3 == 0 {
		res = "Yes"
	}

	fmt.Printf("The number %d is divided by 3 without remainder ? %s", number, res)
}
