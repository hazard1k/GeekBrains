package main

import "fmt"

func main() {
	var amount, percent float32 = 0, 0
	fmt.Print("Enter amount: ")
	fmt.Scan(&amount)
	fmt.Print("\nEnter percent: ")
	fmt.Scan(&percent)
	var inputAmount = amount
	percent /= 100
	for i := 0; i < 5; i++ {
		increase := amount * percent
		amount += increase
		fmt.Printf("for %v year amount increase on %.2f and it will be equal %.2f\n", i+1, increase, amount)
	}

	fmt.Printf("Summary after 5 year amount'll be %.2f\n", amount)
	fmt.Printf("Amount of accrued percents %.2f\n", amount-inputAmount)

}
