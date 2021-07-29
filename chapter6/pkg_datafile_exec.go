package main

import (
	"datafile"
	"fmt"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Println("Сумма:", sum, "\tЭлементов:", sampleCount)
	fmt.Printf("Среднее: %0.2f\n", sum/sampleCount)
}
