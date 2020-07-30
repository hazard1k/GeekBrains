package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const maxChunkSize = 1024

var (
	isInteractive = flag.Bool("i", false, "ask about rewrite existing file")
	isDoRewrite   = flag.Bool("n", false, "do not rewrite existing file")
)

func copyFile(from, to string) error {
	fileFrom, err := os.Open(from)
	if err != nil {
		return err
	}
	defer fileFrom.Close()

	var fileTo *os.File

	if _, err = os.Stat(to); err == nil {
		if *isDoRewrite {
			return nil
		}
		if *isInteractive {
			var doOverride string
			for strings.ToUpper(doOverride) != "Y" {
				fmt.Printf("File %s already exist override? (y/n) ", to)
				fmt.Scan(&doOverride)
				if strings.ToUpper(doOverride) == "N" {
					return nil
				}
			}
		}

		fileTo, err = os.Create(to)
	} else if os.IsNotExist(err) {
		fmt.Println("is not exist")
		fileTo, err = os.Create(to)
	} else {
		return err
	}

	defer fileTo.Close()

	stat, err := fileFrom.Stat()
	if err != nil {
		return err
	}

	fileSize := stat.Size()

	// Цикл по файлу, в зависимости от его частей
	for i := int64(0); i <= fileSize; i += maxChunkSize {
		var chunk int64 = maxChunkSize
		if chunk = maxChunkSize; i+maxChunkSize > fileSize {
			chunk = fileSize - i
		}

		bs := make([]byte, chunk)       // слайс в памяти теперь не занимает весь размер файла
		_, err = fileFrom.ReadAt(bs, i) // читаем частями
		// если есть ошибка чтения и она не конец файла, тогда остановим
		if err != nil && err != io.EOF {
			return err
		}
		fileTo.Write(bs)
	}

	return nil
}

func main() {

	flag.Parse()

	paths := flag.Args()
	source, dest := paths[0], paths[1]

	// Проверим исходный файл
	fi, err := os.Stat(source)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch mode := fi.Mode(); {
	case mode.IsDir(): // Если это дирректория
		fmt.Printf("%v is not a file", fi.Name())
	case mode.IsRegular(): // Если это файл
		fmt.Println("file")
		if err := copyFile(source, dest); err != nil {
			fmt.Printf("Error copy file: %v", err)
		}
	}

}
