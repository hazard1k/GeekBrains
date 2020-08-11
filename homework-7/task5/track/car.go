package track

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const maxSpeed = 400

// Car ..
type Car struct {
	position    int           // Тек. позиция машины
	Speed       int           // Скорость
	Number      int           // Номер на треке
	preparetime time.Duration // Время до готовности
	stopCH      chan struct{}
}

// Stop ..
func (car *Car) Stop() {
	car.stopCH <- struct{}{}
	close(car.stopCH)
}

func (car *Car) prepare() {
	time.Sleep(car.preparetime)
}

// GO ..
func (car *Car) GO(prepare *sync.WaitGroup, finish *sync.WaitGroup, trackCH chan<- *Car, start <-chan struct{}, stop <-chan struct{}) {

	fmt.Printf("\nCAR %v preparing[speed=%v, time to ready=%v]", car.Number, car.Speed, car.preparetime)
	car.prepare()
	fmt.Printf("\nCAR %v prepared", car.Number)
	prepare.Done()
	<-start // Начало гонки
	fmt.Printf("\nCAR %v started", car.Number)
	for {
		select {
		case <-stop:
			finish.Done()
			return
		default:
			car.position += car.Speed
			trackCH <- car
			runtime.Gosched()
		}
		time.Sleep(1 * time.Second)
	}
}

// CarsGenerator ..
func CarsGenerator(count int) []*Car {
	cars := make([]*Car, 0, count)

	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i))
		sp := rand.Intn(maxSpeed)

		car := &Car{
			Speed:       sp,
			preparetime: time.Duration(rand.Intn(10)) * time.Second,
			Number:      i + 1,
		}
		cars = append(cars, car)
	}
	return cars
}
