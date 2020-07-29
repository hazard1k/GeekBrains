package task4

import (
	"fmt"
	"math"
)

// Horse Фигура конь
type Horse struct {
}

// AvailablePoints Возвращает массив точек, в которые конь сможет сделать ход.
func (p *Horse) AvailablePoints(point Point) []Point {
	var res []Point
	for x := -2; x <= 2; x++ {
		for y := -2; y <= 2; y++ {
			// Все шаги в диапазоне -2<=x<=2 -2<=y<=2 если их складывать без учета знака, сумма равняется 3
			if (math.Abs(float64(x)) + math.Abs(float64(y))) == 3 {
				var p *Point
				px := point.X + x
				py := point.Y + y
				if p = New(px, py); p == nil {
					fmt.Printf("Point out of range (x: %v, y: %v)\n", px, py)
					continue
				}

				res = append(res, *p)
				//fmt.Printf("x:%v y:%v\n", hx+x, hy+y)
			}
		}
	}
	return res
}

// Point Точка
type Point struct {
	X, Y int
}

// New инициализация точки
func New(x, y int) *Point {
	if res := isAvailable(x, y); !res {
		return nil
	}
	return &Point{x, y}
}

// isAvailable Проверка на границу доски
func isAvailable(x, y int) bool {
	return !((x < 1 || y < 1) || (x > 8 || y > 8))
}
