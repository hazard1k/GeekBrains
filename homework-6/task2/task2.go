package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

// Image ..
type Image struct {
	*image.RGBA
}

// PrintLine ..
func (i *Image) PrintLine(x1, y1, x2, y2 int, color color.Color) {
	y1 = i.Rect.Size().Y - y1
	y2 = i.Rect.Size().Y - y2

	var k float32 // при x1 = x2 угловой коэф = 0
	if x1 != x2 {
		k = float32((y2 - y1) / (x2 - x1)) // Угловой коэфициент
	}

	b := float32(y1) - k*float32(x1)
	var from, to int
	var byx bool
	if y1 == y2 {
		// гризонтальная прямая цикл по оси x
		from, to, byx = x1, x2, true
	} else {
		from, to = y2, y1
	}

	for x := from; x <= to; x++ {
		if byx {
			y := k*float32(x) + b
			i.Set(int(x), int(y), color)
		} else {
			y := x1
			i.Set(int(y), int(x), color)
		}
	}
}

// PrintHorLine ..
func (i *Image) PrintHorLine(x1, x2, y int, color color.Color) {
	i.PrintLine(x1, y, x2, y, color)
}

// PrintVertLine ..
func (i *Image) PrintVertLine(y1, y2, x int, color color.Color) {
	i.PrintLine(x, y1, x, y2, color)
}

func main() {
	green := color.RGBA{0, 255, 0, 255}

	canvas := &Image{image.NewRGBA(image.Rect(0, 0, 200, 201))}
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	// Нарисуем прямоугольник
	canvas.PrintHorLine(50, 150, 20, color.RGBA{255, 0, 0, 255})
	canvas.PrintHorLine(50, 150, 180, color.RGBA{255, 0, 0, 255})
	canvas.PrintVertLine(20, 180, 50, color.RGBA{255, 0, 0, 255})
	canvas.PrintVertLine(20, 180, 150, color.RGBA{255, 0, 0, 255})
	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, canvas)
}
