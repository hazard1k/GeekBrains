package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

/*
	Внести в телефонный справочник дополнительные данные.
	Реализовать сохранение json-файла на диске с помощью пакета ioutil при изменении данных.
*/

// PhoneBook Справочник
type PhoneBook struct {
	Phones map[string]string
}

// New Инициализация телефонной книги
func New() *PhoneBook {
	return &PhoneBook{Phones: make(map[string]string, 0)}
}

// Add Добавление номера телефона по ключу name
func (p *PhoneBook) Add(phone, name string) {
	p.Phones[phone] = name
}

// Del Удаление телефонного номера
func (p *PhoneBook) Del(phone string) {
	delete(p.Phones, phone)
}

// Get Получение имени по телефону
func (p *PhoneBook) Get(phone string) (string, bool) {
	var name string
	var ok bool
	name, ok = p.Phones[phone]
	return name, ok
}

// List Список телефонов
func (p *PhoneBook) List() {
	for phone, name := range p.Phones {
		fmt.Printf("%s - %s\n", phone, name)
	}
}

// SaveJSON Сохранение справочника
func (p *PhoneBook) SaveJSON() error {
	var bookJSON []byte
	bookJSON, err := json.Marshal(p.Phones)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("book.json", bookJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(temp))
	return nil
}
func main() {
	book := New()
	book.Add("+70000000000", "Харрисон Форд")
	book.Add("+70000000001", "Вин Дизель")
	book.Add("+70000000002", "Джордж Клуни")
	book.SaveJSON()
}
