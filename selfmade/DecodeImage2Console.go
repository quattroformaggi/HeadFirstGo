// REF: https://golangdocs.com/golang-image-processing
// Декодирование изображений в консоль, Console Image Decoding
package main

import (
	"bufio"
	"fmt"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	sampleFile, err := os.Open("decode_kiryu.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer sampleFile.Close() //закрыть файл по завершении

	// рекомендуется метод image.Decode, может понять и декодировать любой формат
	img, err := jpeg.Decode(sampleFile)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(img) не нужен

	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y / 51 // 51 * 5 = 255
			if level == 5 {
				level--
			}
			fmt.Print(levels[level])
		}
		fmt.Print("\n")
	}

	//ожидание ввода
	fmt.Print("Нажмите любую кнопку для выхода: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())
}
