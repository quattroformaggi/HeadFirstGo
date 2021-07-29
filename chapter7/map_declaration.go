package main

import (
	"fmt"
)

func main() {
	var ranks map[string]int = make(map[string]int)
	ranks["GOLD"] = 1
	ranks["SILVER"] = 2
	ranks["BRONZE"] = 3
	fmt.Println("Призовые места:", ranks) //в алфавитном порядке
	for strings, ints := range ranks {
		fmt.Println(strings, ints) //обычный порядок
	}
	constants := map[string]float64{"Pi": 3.14, "g": 9.8, "e": 2.7}
	fmt.Println("Константы:", constants)
}
