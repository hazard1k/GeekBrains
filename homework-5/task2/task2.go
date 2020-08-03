package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const maxChunkSize = 1048576 // 1mb, размер части, по которой считывать файл

// вариант чтения по частям
func own() {

	file, err := os.Open("fileread.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// getting size of file
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fileSize := stat.Size()

	// Цикл по файлу, в зависимости от его частей
	for i := int64(0); i <= fileSize; i += maxChunkSize {
		bs := make([]byte, maxChunkSize) // слайс в памяти теперь не занимает весь размер файла
		_, err = file.ReadAt(bs, i)      // читаем частями
		// если есть ошибка чтения и она не конец файла, тогда остановим
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
		//fmt.Print(string(bs))
	}

}

// Изначальный вариант
func gb() {

	file, err := os.Open("fileread.txt")
	if err != nil {
		return
	}
	defer file.Close()

	// getting size of file
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// reading file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

}
func main() {
	start := time.Now()
	own()
	fmt.Println(time.Since(start))
	start = time.Now()
	gb()
	fmt.Println(time.Since(start))
	// Время идентичное, а памяти использует меньше
}
