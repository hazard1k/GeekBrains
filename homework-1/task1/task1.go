package main

import "fmt"

const usd = 71.29

func main() {
	var rurAmount float32
	fmt.Printf("1 USD = %v RUR\n", usd)
	fmt.Print("Enter amount(RUR): ")
	fmt.Scan(&rurAmount)
	fmt.Printf("%.2f RUR = %.2f USD", rurAmount, rurAmount/usd)
}
