package main

import "fmt"

/*
	Реализовать очередь.
	Это структура данных, работающая по принципу FIFO (First Input — first output, или «первым зашел — первым вышел»).
*/

// Queue сама очередь
// Через двунаправленный список
type Queue struct {
	head   *Node // Голова очереди
	tail   *Node // Хвост очереди
	length int   // Длина очереди
}

// Node элемент очереди
type Node struct {
	prev  *Node       // Указатель на предыдущего
	next  *Node       // Указатель на следующего
	value interface{} // Значение элемента
}

// New Создание очереди
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// Enqueue Добавление элемента в очередь
func (q *Queue) Enqueue(value interface{}) {

	node := &Node{value: value} // Создадим новый элемент  очереди
	if q.length == 0 {
		// Если очередь была пуста(первый добавляемый элемент), то указатели головы и хвоста должны ссылаться на него
		q.head, q.tail = node, node
	} else {
		// Если это не первый элемент, то
		node.prev = q.tail // в предыдущий элемент ноды запишем указатели на новую ноду
		q.tail.next = node
		q.tail = node // и сделаем новый элемент хвостом очереди
	}
	q.length++ // увеличим размер очереди
}

// Dequeue Извлечение элемента из очереди
func (q *Queue) Dequeue() interface{} {
	if q.length == 0 {
		return nil
	}
	node := q.head     // заберем элемент с головы очереди
	q.head = node.next // И поменяем голову на след элемент
	q.length--         // уменьшим размер очереди
	return node.value
}

//================================Вторая реализация

// Queue2 Реализация2
// Через обычный слайс
type Queue2 struct {
	Queue []interface{}
}

// New2 Создание очереди второй реализации
func New2() *Queue2 {
	return &Queue2{}
}

// Enqueue Добавление элемента в очередь
func (q *Queue2) Enqueue(value interface{}) {
	q.Queue = append(q.Queue, value) // Просто добавим в слайс новое значение
}

// Dequeue Извлечение элемента из очереди
func (q *Queue2) Dequeue() interface{} {
	if len(q.Queue) == 0 {
		return nil
	}
	value := q.Queue[0]   // возьмем первое значение
	q.Queue = q.Queue[1:] // Удалим первое значение из очереди
	return value
}

func main() {

	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println("Результат первой реализации")
	for val := q.Dequeue(); val != nil; val = q.Dequeue() {
		fmt.Println(val)
	}

	q2 := New2()
	q2.Enqueue(1)
	q2.Enqueue(2)
	q2.Enqueue(3)
	q2.Enqueue(4)
	q2.Enqueue(5)
	q2.Enqueue(6)
	fmt.Println("Результат второй реализации")

	for val := q2.Dequeue(); val != nil; val = q2.Dequeue() {
		fmt.Println(val)
	}
}
