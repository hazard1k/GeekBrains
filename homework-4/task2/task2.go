package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// phoneBookRecord структура записи
type phoneBookRecord struct {
	phoneNumber string
	name        string
}

// phoneBook сама телефонная книга
type phoneBook []phoneBookRecord

// prepareNumber  функция парсинга телефона
func prepareNumber(number string) int64 {
	if number[0] == '+' {
		number = strings.TrimLeft(number, "+")
	}
	if strings.Contains(number, "(") {
		number = strings.ReplaceAll(number, "(", "")
	}
	if strings.Contains(number, ")") {
		number = strings.ReplaceAll(number, ")", "")
	}

	num, _ := strconv.ParseInt(number, 10, 64)

	return num
}

// Релизация интерфейса Sort
func (phb phoneBook) Len() int {
	return len(phb)
}

func (phb phoneBook) Less(i, j int) bool {
	return prepareNumber(phb[i].phoneNumber) < prepareNumber(phb[j].phoneNumber)
}

func (phb phoneBook) Swap(i, j int) {
	phb[i], phb[j] = phb[j], phb[i]
}

func main() {
	book := &phoneBook{{"+79990000000", "People 5"}, {"+7(850)0000000", "People 4"}, {"+72000000000", "People 2"}, {"+7(532)0000000", "People 3"}, {"+70010000000", "People 1"}}
	fmt.Println(*book)
	sort.Sort(book)
	fmt.Println(*book)
}
