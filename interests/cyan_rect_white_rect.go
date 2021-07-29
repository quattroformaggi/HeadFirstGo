package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

//https://golang-blog.blogspot.com/2020/05/create-image-in-golang.html
func main() {
	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Цвета определяются значениями
	// Red, Green, Blue, Alpha uint8.
	cyan := color.RGBA{100, 200, 200, 0xff}

	// Установить цвет для каждого пикселя.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			// верхний левый квадрант
			case x < width/2 && y < height/2:
				img.Set(x, y, cyan)
			// нижний правый квадрант
			case x >= width/2 && y >= height/2:
				img.Set(x, y, color.White)
			default:
				// Использовать нулевое значение.
			}
		}
	}

	// Кодировать как PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
