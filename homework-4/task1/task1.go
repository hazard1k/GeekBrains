package main

import "fmt"

type startable interface {
	start()
	stop()
}

type engine struct {
	fuel   string
	volume float32
}
type truck struct {
	model string
	engine
}

func (t *truck) start() {
	fmt.Printf("Starting truck engine fuel: %s, volume: %f", t.fuel, t.volume)
}
func (t *truck) stop() {
	fmt.Println("Stopping truck engine")
}

type washer struct {
	name    string
	payload float32
}

func (w *washer) start() {
	fmt.Printf("Starting washer machine name: %s with payload: %f", w.name, w.payload)
}
func (w *washer) stop() {
	fmt.Println("Stopping washer machine")
}

func startEngine(eng startable) {
	eng.start()
}
func stopEngine(eng startable) {
	eng.stop()
}

func main() {
	kitchenWasher := &washer{name: "Indesit", payload: 5.2}
	volvoTruck := &truck{engine: engine{fuel: "Diesel", volume: 12.002}, model: "FH"}
	startEngine(kitchenWasher)
	stopEngine(kitchenWasher)
	startEngine(volvoTruck)
	stopEngine(volvoTruck)
}
