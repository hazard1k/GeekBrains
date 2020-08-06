package track

import (
	"fmt"
	"sync"
)

const maxCars = 10

type trackChannel *Car

// Track ..
type Track struct {
	cars      []*Car
	length    int
	trackCH   chan *Car // Канал для сообщения позиции авто на треке
	winnersCH chan *Car // Канал для вывода победителей
}

// StartRace ..
func (t *Track) StartRace() {
	t.winnersCH = make(chan *Car, len(t.cars))
	endRaceCH := make(chan struct{})
	// Функция для мониторинга позиций машин
	go func() {
		for {
			select {
			case car := <-t.trackCH:
				if car.position >= t.length {
					// это финиш авто
					t.winnersCH <- car       // в список финишировавших
					car.stopCH <- struct{}{} // остановим поток с машиной
					//car.Stop()
				}
				fmt.Printf("\nCAR %v cur position %v/%v", car.Number, car.position, t.length)
			case <-endRaceCH: // Для завершения текущей горутины
				close(t.winnersCH) // Закроем канал победителей, чтобы программа завершилась корректно
				return
			default:
			}
		}
	}()
	// Две группы одна ждет подготовки авто, вторая завершения всех горутин(финиш)
	wgPrepare, wgFinish := &sync.WaitGroup{}, &sync.WaitGroup{}
	startCH := make(chan struct{})
	for _, car := range t.cars {
		wgFinish.Add(1)
		wgPrepare.Add(1)
		// на каждую горутину создадим свой завершающий канал, чтобы останавливать их по мере финиширования
		stopCH := make(chan struct{})
		car.stopCH = stopCH
		go car.GO(wgPrepare, wgFinish, t.trackCH, startCH, stopCH)
	}
	wgPrepare.Wait() // Ждем готовности
	fmt.Print("\n================\n")
	fmt.Print("Все готовы")
	fmt.Println("\n================")
	close(startCH)          // Синхронизация всех горутин
	wgFinish.Wait()         // Ждем всех, пока не доедут(завершение горутин)
	endRaceCH <- struct{}{} // Закроем горутину, которая слушала авто
	fmt.Println("\n================")
	fmt.Print("Все завершили гонку")
	fmt.Println("\n================")
	fmt.Println("\n\n Победители:")
	var placeCnt int = 1
	// Цикл по накопленным приехавшим машинам
	for winner := range t.winnersCH {
		fmt.Printf("%v) Car number - %v, speed: %v\n", placeCnt, winner.Number, winner.Speed)
		placeCnt++
	}
}

// RegisterCars Инициализирует все машины необходимыми каналами
func (t *Track) RegisterCars(cars []*Car) {
	t.cars = cars
}

// New ..
func New() *Track {

	return &Track{
		cars:    nil,
		length:  1000,
		trackCH: make(chan *Car),
		//start: make(chan struct{}),
		//end:   make(chan struct{}),

	}
}
