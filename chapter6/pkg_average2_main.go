//average2 conducts an average sum of numbers
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func Avg(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	args := os.Args[1:]
	var numbers []float64
	for _, arg := range args {
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	fmt.Printf("Среднее:%0.2f\n", Avg(numbers...))
}
