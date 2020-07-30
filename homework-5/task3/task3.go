package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// User ..
type User struct {
	ID    int
	Name  string
	Phone string
}

// CSVFormat ..
func (u *User) CSVFormat() []string {
	return []string{strconv.Itoa(u.ID), u.Name, u.Phone}
}

func csvReader() (users []User) {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error openning file:", err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	r.Comma = '|'
	for {
		record, e := r.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		userID, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println("Not numeric UserID")
		}
		user := User{userID, record[1], record[2]}
		users = append(users, user)

	}
	return users
}

func csvWriter() {
	file, err := os.OpenFile("data.csv", os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error openning file:", err)
	}
	defer file.Close()
	users := []User{User{1, "Vasya", "89009990099"}, User{2, "Petya", "89009990088"}, User{3, "Alex", "89009990011"}}
	w := csv.NewWriter(file)
	w.Comma = '|'
	w.UseCRLF = true
	for _, user := range users {
		if err := w.Write(user.CSVFormat()); err != nil {
			fmt.Println("Error writing to file:", err)
		}

	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	csvWriter()
	fmt.Println(csvReader())
}
