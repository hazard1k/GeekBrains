package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getWalkFunc(what string) func(path string, info os.FileInfo, err error) error {

	return func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			absPath, _ := filepath.Abs(path)
			file, err := os.Open(absPath)
			if err != nil {
				return err
			}
			var lineCounter int
			reader := bufio.NewReader(file)
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						break
					} else {
						return err
					}
				}
				lineCounter++

				if idx := strings.Index(strings.ToUpper(line), strings.ToUpper(what)); idx != -1 {
					fmt.Printf("Found at %s (%v,%v)\n", absPath, lineCounter, idx+1)
				}
			}
		}

		return nil
	}
}

func main() {
	var where string = "./" // Где искать, по дефолту тек. папка
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Set the what to find")
		return
	} else if len(args) == 2 {
		where = args[1]
	}
	what := args[0] // Что искать
	_ = what
	if err := filepath.Walk(where, getWalkFunc(what)); err != nil {
		log.Fatal(err)
	}

}
