/*Вывод файла в консоль без разметки, требуется 'notes.txt'
File output in console w/o markup, requires 'notes.txt'*/

package main

import (
	"fmt"
	"log"
	"os"
)

func checkFileExistence(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//As of Go 1.16, ioutil.ReadFile simply calls os.ReadFile. (c) pkg
	fileData, err := os.ReadFile("notes.txt") //считывает файл по заданному имени
	checkFileExistence(err)                   //проверяет, успешно ли считывание, присутствует ли файл
	strData := string(fileData)               //переводит в строковый тип данные из файла
	fmt.Println("[notes.txt]\n", strData)     //выводит эти данные как они представлены в файле
}
