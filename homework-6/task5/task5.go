package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

// BoardSettings ..
type BoardSettings struct {
	heightPx, widthPx, colsCnt, rowsCnt int
}

// ChessBoard ..
type ChessBoard struct {
	Settings BoardSettings
}

// New ..
func New(bs BoardSettings) *ChessBoard {
	return &ChessBoard{bs}
}

// RenderBoard ..
func (cb *ChessBoard) RenderBoard() *image.RGBA {
	// Создадим сам "холст"
	containerImage := image.NewRGBA(image.Rect(0, 0, cb.Settings.heightPx, cb.Settings.widthPx))
	containerColor := color.RGBA{0, 100, 0, 255} //  R, G, B, Alpha

	draw.Draw(containerImage, containerImage.Bounds(), &image.Uniform{containerColor}, image.ZP, draw.Src)
	fieldHeight := cb.Settings.heightPx / cb.Settings.colsCnt
	fieldWidth := cb.Settings.widthPx / cb.Settings.rowsCnt
	for i := 0; i < cb.Settings.rowsCnt; i++ {
		for j := 0; j < cb.Settings.colsCnt; j++ {
			x1 := i * fieldHeight
			y1 := j * fieldWidth
			x2 := x1 + fieldHeight
			y2 := y1 + fieldWidth
			fieldRect := image.Rect(x1, y1, x2, y2) // Создадим само поле
			fieldColor := color.RGBA{237, 232, 232, 255}
			// поменяем цвет поля
			if (i+j)%2 != 0 {
				fieldColor = color.RGBA{18, 16, 16, 255}
			}
			// нарисуем поле на "холсте"
			draw.Draw(containerImage, fieldRect, &image.Uniform{fieldColor}, image.ZP, draw.Src)
		}
	}

	return containerImage
}

func main() {
	chessBoard := New(BoardSettings{800, 800, 8, 8})
	boardImg := chessBoard.RenderBoard()

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, boardImg)
}
