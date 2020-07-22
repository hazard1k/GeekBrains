package main

import "fmt"

/*
Инициализировать несколько экземпляров структур. Применить к ним различные действия.
Вывести значения свойств экземпляров в консоль.
*/

// Vehicle ..
type Vehicle struct {
	brand         string //Марка
	year          string //Год
	volume        int    //Объем багажника/кузова
	isEngStarted  bool   //запущен ли двигатель
	isOpenWindows bool   //открыты ли окна
	vulumeload    int    //насколько заполнен объем багажника
}

type car struct {
	*Vehicle
	gearType string
}

type truck struct {
	Vehicle
	maxPayload float32
}

func main() {
	volvoTruck := truck{Vehicle: Vehicle{brand: "Volvo"}, maxPayload: 20000}
	fmt.Println(volvoTruck.brand, volvoTruck.maxPayload)
	volvoTruck.maxPayload += 2000
	fmt.Println(volvoTruck.brand, volvoTruck.maxPayload)
	golfCar := car{Vehicle: &Vehicle{brand: "Volkswagen", isOpenWindows: true, isEngStarted: false}}
	fmt.Println(golfCar.isEngStarted)
	golfCar.isEngStarted = !golfCar.isEngStarted
	fmt.Println(golfCar.isEngStarted)
}
