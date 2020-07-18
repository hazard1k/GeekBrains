package main

import (
	"fmt"
	"math/big"
)

func main() {

	// Тут не хватает размера uint64 на последних значениях ряда
	// var n1Num, n2Num, cur uint64 = 0, 1, 0
	// for i := 0; i < 100; i++ {
	// 	if i == 1 {
	// 		cur++
	// 	} else {
	// 		n2Num = n1Num
	// 		n1Num = cur
	// 		cur = n1Num + n2Num
	// 	}
	// 	fmt.Printf("%d) %v\n", i+1, cur)
	// }

	// Поэтому через big

	var n2Num, n1Num, cur = big.NewInt(0), big.NewInt(1), big.NewInt(0)
	for i := 0; i < 100; i++ {
		if i < 2 {
			fmt.Println(i)
		} else {
			cur.Add(n1Num, n2Num)
			n2Num.SetString(n1Num.String(), 10)
			n1Num.SetString(cur.String(), 10)
			fmt.Println(cur)
		}
	}
}
