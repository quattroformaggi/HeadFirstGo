package main

import (
	"errors"
	"fmt"
	"math"
)

//task 5,6,false,$$$$,11,30,true,hahaha
func main() {
	fmt.Println("Hello World!")
	fmt.Printf("%v", "\na")      // a
	fmt.Printf("%#v", "\na")     // "\na"
	fmt.Printf("%f литров", 1.9) //1.900000 литров
	//% ширина-числа.ширина-др-части тип-глагола-формат
	fmt.Printf("%5.3f", 11.1337) //11.134
	sayHello()
	sayTimesX("man...", 5)
	fmt.Println(sayAndReturn("fisting is 300$"))
	fmt.Printf("%6.2f", circleArea(-9999999.9999999))
	err := errors.New("Ошиб очка!")
	fmt.Println(err.Error())
	fmt.Println(err)
	outFloat, outString, outInteger := returns()
	fmt.Println("Float:", outFloat, "String:", outString, "Integer:", outInteger)
}

func sayHello() { //функция приветствия
	fmt.Println("PC говорит \"Приветствую!\"")
}

func sayTimesX(say string, X int) { //функция 'сказать X раз'
	for i := 0; i < X; i++ {
		fmt.Println(say)
	}
}

func sayAndReturn(say string) string { //функция 'сказать и вернуть'
	fmt.Println(say, " и возвращаю...")
	return say
}

func circleArea(rad float64) float64 {
	return math.Pow(rad, 2) * math.Pi
}

func returns() (float64, string, int) {
	return 13.37, "l33t", 2021
}
