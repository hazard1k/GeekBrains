package main

/*
Описать несколько структур — любой легковой автомобиль и грузовик. Структуры должны
содержать марку авто, год выпуска, объем багажника/кузова, запущен ли двигатель, открыты
ли окна, насколько заполнен объем багажника.
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

// Car ..
type Car struct {
	common   Vehicle
	gearType string
}

// Truck ..
type Truck struct {
	common     Vehicle
	maxPayload float32
}

func main() {

}
