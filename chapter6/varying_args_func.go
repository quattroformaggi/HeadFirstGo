package main

import "fmt"

func Sum(numbers ...int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func Avg(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	fmt.Println("Sum:", Sum(9, 7, 5), "\tAvg:", Avg(9, 7, 5))
}
//magnets etc