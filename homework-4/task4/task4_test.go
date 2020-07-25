package task4

import (
	"fmt"
	"reflect"
	"testing"
)

var tests = []struct {
	x   int
	y   int
	exp []Point
}{
	{1, 1, []Point{Point{2, 3}, Point{3, 2}}},
	{8, 6, []Point{Point{6, 5}, Point{6, 7}, Point{7, 4}, Point{7, 8}}},
	{3, 2, []Point{Point{1, 1}, Point{1, 3}, Point{2, 4}, Point{4, 4}, Point{5, 1}, Point{5, 3}}},
	{5, 5, []Point{Point{3, 4}, Point{3, 6}, Point{4, 3}, Point{4, 7}, Point{6, 3}, Point{6, 7}, Point{7, 4}, Point{7, 6}}},
}

func TestCreatePoint(t *testing.T) {
	fmt.Println(t.Name())
	p := New()
	if p == nil {
		t.Errorf("The point must be created")
	}
}

//TODO сделать сравнение слайсов для вывода ошибок
func TestHorseSteps(t *testing.T) {
	fmt.Println(t.Name())
	for _, e := range tests {
		points := AvailablePoints(e.x, e.y)
		if !reflect.DeepEqual(points, e.exp) {
			t.Errorf("Find(%v, %d) = %v, expected %d", e.x, e.y, points, e.exp)
		}
	}
}
