package main

import (
	"GeekBrains/homework-7/task5/track"
)

func main() {
	trackGame := track.New()

	trackGame.RegisterCars(track.CarsGenerator(10))
	trackGame.StartRace()
}
