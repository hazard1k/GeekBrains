package main

import (
	"fmt"
)

const countNum = 10

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; x < countNum; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // канал закрыт и пуст
			}
			squares <- x * x
		}
		close(squares)
	}()

	// печать
	for {
		num, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(num)
	}
}
