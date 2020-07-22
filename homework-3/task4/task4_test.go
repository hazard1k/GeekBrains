package main

import (
	"testing"
)

func TestCreateBook(t *testing.T) {
	book := New()
	if book == nil {
		t.Errorf("Book must be initialized")
	}

}

func TestBookAddGetPhone(t *testing.T) {
	book := New()
	book.Add("+70000000000", "Харрисон Форд")
	name, ok := book.Get("+70000000000")
	if !ok {
		t.Errorf("В книге должен находиться номер")
	}
	if name != "Харрисон Форд" {
		t.Errorf("Имя должно быть Харрисон Форд")
	}

}

func TestBookDelPhone(t *testing.T) {
	book := New()
	book.Add("+70000000000", "Харрисон Форд")
	book.Del("+70000000001")
	_, ok := book.Get("+70000000001")
	if ok {
		t.Errorf("Номер телефона должен был быть удален")
	}

}
