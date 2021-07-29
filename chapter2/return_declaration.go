package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите оценку:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Введено:", input)
}
