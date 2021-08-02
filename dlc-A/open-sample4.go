//DLC-A
package main

import (
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
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("sample-file.txt", options, os.FileMode(1800))
	check(err)
	_, err = file.Write([]byte("a sample text!\n"))
	check(err)
	err = file.Close()
	check(err)
	fmt.Println("Файл создан | обновлён!")
}
