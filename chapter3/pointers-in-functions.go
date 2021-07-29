package main

import (
	"fmt"
)

func main() {
	amount := 6
	double(amount)
	fmt.Println(amount, "по адресу", &amount, "\tНЕ изменился...")
	double2(&amount)
	fmt.Println(amount, "по адресу", &amount, "\tизменился!")
}

func double(num int) {
	num *= 2
}
func double2(num *int) {
	*num *= 2
}
