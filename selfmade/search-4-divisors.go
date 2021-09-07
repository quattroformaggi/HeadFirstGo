//Поиск всех делителей заданного числа.(5min)
package main

import (
	"fmt"
	"os"
)

func main() {
	var input int
	fmt.Print("\nВведите число:")
	fmt.Scan(&input)
	if input < 0 {
		input *= -1
	} else if input == 0 || input == 1 || input == -1 {
		fmt.Println(input, " - тривиальный делитель самого себя")
		os.Exit(1)
	}
	fmt.Println("\nВозможные делители числа", input, ":")
	for i := 1; i <= input; i++ {
		if input%i == 0 {
			fmt.Print(i, ";")
		}
	}
}
