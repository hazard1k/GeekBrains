package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func mirroredQuery() string {
	var hostnames = []string{
		"https://yahoo.com",
		"https://rambler.ru",
		"https://google.com",
		"https://ya.ru",
	}
	responses := make(chan string, len(hostnames))
	defer func() { close(responses) }()
	wg := &sync.WaitGroup{}
	for _, host := range hostnames {
		wg.Add(1)
		go func(hostname string) {
			start := time.Now()
			_ = request(hostname)
			end := time.Now().Sub(start)
			fmt.Printf("get %s, took %s\n", hostname, end)
			responses <- hostname
			wg.Done()
		}(host)
	}
	wg.Wait()
	return <-responses // возврат самого быстрого ответа
}

func request(hostname string) string {

	resp, err := http.Get(hostname)
	if err != nil {
		log.Printf("Error get %s resourse : %s", hostname, err.Error())
		return ""
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%s", err)
		return ""
	}
	return string(contents)

}

func main() {
	first := mirroredQuery()
	fmt.Println("The first was", first)
}
