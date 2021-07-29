package main

import (
	"fmt"
)

func main() {
	var price int = 100
	fmt.Println("Цена равна", price, "р.")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Налоги составляют", tax, "р.")

	var total float64 = float64(price) + tax
	fmt.Println("Полная стоимость равна", total, "р.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "р. доступно.")
	fmt.Println("Из бюджета?", total <= float64(availableFunds))
}