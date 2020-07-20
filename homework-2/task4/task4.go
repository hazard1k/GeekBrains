package main

import "fmt"

/*
Заполнить массив из 100 элементов различными простыми числами. Натуральное число, которое больше единицы, называется простым, если оно делится только на себя и на единицу. Для нахождения всех простых чисел не больше заданного числа n, следуя методу Эратосфена, нужно выполнить следующие шаги:
Выписать подряд все целые числа от двух до n (2, 3, 4, ..., n).
Пусть переменная p изначально равна 2 — первому простому числу.
Зачеркнуть в списке числа от 2p до n, считая шагами по p (это будут числа, кратные p: 2p, 3p, 4p, ...).
Найти первое не зачеркнутое число в списке, превышающее p, и присвоить значению переменной p это число.
Повторять шаги c и d, пока возможно.

*/
func main() {
	n := 550 //этого будет достаточно для получения 100 простых чисел
	// Массив для "зачеркивания", правда наоборот, чтобы не делать цикл инициализации в true
	var notNaturalNumber = make([]bool, n)
	// Результирующий массив найденных натуральных чисел
	var simpleNumbers = []int{}

	for i := 2; i < n; i++ { //цикл по всем целым числам, ограничиваемся длиной созданного слайса
		// Если "зачеркнуто" - пропускаем
		if notNaturalNumber[i] {
			continue
		}
		// Сам цикл зачеркивания
		for j := 2; j < n; j++ {
			// Условие чтобы не выйти за размер слайса
			if i*j > n-1 {
				break
			}
			// "вычеркиваем" число
			notNaturalNumber[i*j] = true
		}
		// добавляем в результирующий слайс простых чисел
		simpleNumbers = append(simpleNumbers, i)
	}
	// Выводим 100 чисел из слайса
	// При длине ряда 550 целых чисел, находится 101 простое, поэтому обрежем(ну и так чтобы попробовать)
	fmt.Println(simpleNumbers[:100])
}

//521 - 98 число
