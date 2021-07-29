package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Открыт", fileName)
	return os.Open(fileName)
}
func CloseFile(file *os.File) {
	fmt.Println("Закрыт файл")
	file.Close()
}
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file)
	scnr := bufio.NewScanner(file)
	for scnr.Scan() {
		number, err := strconv.ParseFloat(scnr.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	if scnr.Err() != nil {
		return nil, scnr.Err()
	}
	return numbers, nil
}

func main() {
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Сумма: %0.2f\n", sum)
}
