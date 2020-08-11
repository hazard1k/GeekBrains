package main

import (
	"fmt"
	"time"
)

// Реализация через тикер
const spinnerTimeSec = 2

func spinner(delay time.Duration, cancel <-chan int) {
	for {
		for _, r := range "-\\|/" {
			select {
			case <-cancel:
				return
			default:
				fmt.Printf("%c\r", r)
				time.Sleep(delay)
			}
		}
	}
}

func main() {

	tick, cancel := make(<-chan time.Time), make(chan int) // Создадим канал для тиком и для отмены спинера
	tick = time.Tick(1 * time.Second)                      // создаем поток секундных «тиков»
	go spinner(50*time.Millisecond, cancel)
	for countdown := spinnerTimeSec; countdown > 0; countdown-- {
		<-tick //ждем секунду
	}
	cancel <- 1 // отправляем отмуну спинеру
	close(cancel)

}
