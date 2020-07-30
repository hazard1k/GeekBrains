package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	now := time.Now()
	fmt.Println(now.Format("_2.1.06 .000 ms"))

	dateStr := "30.12.1986 11:49"
	date, _ := time.Parse("02.01.2006 15:04", dateStr)
	fmt.Println(date.Format("02.01.2006 15:04"))

	duration := time.Since(start)
	fmt.Printf("Operation took: %s", duration)
}
