package task4

import (
	"errors"
	"fmt"
	"math"
)

// Point Точка
type Point struct {
	x, y int
}

// New инициализация точки
func New() *Point {
	return &Point{}
}

// SetXY Создает точку
func (p *Point) SetXY(x, y int) error {
	if (x < 1 || y < 1) || (x > 8 || y > 8) {
		return errors.New("coordinates out of range")
	}
	p.x, p.y = x, y
	return nil
}

// AvailablePoints Возвращает массив точек, в которые конь сможет сделать ход.
func AvailablePoints(hx, hy int) []Point {
	var res []Point
	for x := -2; x <= 2; x++ {
		for y := -2; y <= 2; y++ {
			// Все шаги в диапазоне -2<=x<=2 -2<=y<=2 если их складывать без учета знака, сумма равняется 3
			if (math.Abs(float64(x)) + math.Abs(float64(y))) == 3 {
				p := Point{}
				px := hx + x
				py := hy + y
				if err := p.SetXY(px, py); err != nil {
					fmt.Printf("Point %v (x: %v, y: %v)\n", err, px, py)
					continue
				}
				res = append(res, p)
				//fmt.Printf("x:%v y:%v\n", hx+x, hy+y)
			}
		}
	}
	return res
}
