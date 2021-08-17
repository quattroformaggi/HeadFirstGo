//создание RGB дорожки размера N*3
package main

import (
	//"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func checkErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() { //.Pix[N] 0-R,1-G,2-B,3-A
	var xsize int = 24 //размеры полоски
	var ysize int = 3
	testImg := image.NewRGBA(image.Rect(0, 0, xsize, ysize)) //создание пустой картинки в буфере
	for R := 0; R < xsize*4; R += 4 {                        //черчение красным
		testImg.Pix[R] = 255
		testImg.Pix[R+3] = 255 //и замена альфа канала
	}
	for G := xsize*4 + 1; G < xsize*8; G += 4 { //зелёным
		testImg.Pix[G] = 255
		testImg.Pix[G+2] = 255
	}
	for B := xsize*4*2 + 2; B < xsize*12; B += 4 { //синим
		testImg.Pix[B] = 255
		testImg.Pix[B+1] = 255
	}

	output, err := os.Create("testImage.png") //выводит буфер в файл
	png.Encode(output, testImg)               //требует интерфейс записи и интерфейс картинки
	output.Close()                            //закрытие вывода
	checkErrors(err)                          //проверка на наличие ошибок
	// fmt.Println(testImg.Pix) //выводит коды пикселей
	// fmt.Println("Успешно выполнено.")
}
