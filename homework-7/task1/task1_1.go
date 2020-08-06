package task1

import (
	"fmt"
	"time"
)

// Реализация через таймер, но тут необходим fmt.Scan(&a)
const spinnerTimeSec = 2

func spinner(delay time.Duration, cancel <-chan time.Time) {
	for {
		for _, r := range "-\\|/" {
			select {
			case <-cancel:
				fmt.Print(" \r")
				return
			default:
				fmt.Printf("%c\r", r)
				time.Sleep(delay)
			}
		}
	}
}

func main() {
	go spinner(50*time.Millisecond, time.NewTimer(spinnerTimeSec*time.Second).C)
	var a int32
	fmt.Scan(&a)
}
