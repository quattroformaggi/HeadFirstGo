/*Вывод файла в консоль построчно, требуется 'lines.txt'
Line-by-line file output in console, requires 'lines.txt'*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func check(err error) { //проверка ошибок
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	filePtr := flag.String("filepath", "lines.txt", "file path to read from") //создание флага для считывания файла
	flag.Parse()                                                              //разбирает аргументы командной строки

	file, err := os.Open(*filePtr) //открываем файл используя флаг
	check(err)                     //проверяем успешность открытия файла
	defer func() {                 //отложенный вызов анонимной(!) функции
		if err = file.Close(); err != nil { //если файл закрыт или любая другая ошибка
			log.Fatal(err) //вывести сообщение об ошибке
		}
	}()

	scanner := bufio.NewScanner(file) //создаём сканер для считывания файла
	var lineNum = 0                   //создаём переменную номера строки как в блокноте
	for scanner.Scan() {              //пока сканер считывает строки
		lineNum++                                  //увеличиваем счётчик номера строки
		fmt.Println(lineNum, "\t", scanner.Text()) //выводим его вместе с текстом
	}
	check(scanner.Err()) //проверяем, успешно ли прошло считывание
}
