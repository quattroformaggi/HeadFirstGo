package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//ввод
	fmt.Print("Введите оценку 1..100:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	//условность
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "проходной балл достигнут"
	} else {
		status = "провалено..."
	}
	fmt.Println("Оценка, равная", grade, "--", status)
}
