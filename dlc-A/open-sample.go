//DLC-A
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.OpenFile("sample-file.txt", os.O_RDONLY, os.FileMode(0600))
	check(err)
	defer file.Close() //отложенное закрытие
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) //считывает нормально
	}
	check(scanner.Err())
}
