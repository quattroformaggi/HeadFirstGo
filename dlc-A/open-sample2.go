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
	file, err := os.OpenFile("sample-file.txt", os.O_WRONLY, os.FileMode(1800))
	check(err)
	_, err = file.Write([]byte("a sample text!\n")) //перезаписывает файл поверх того, что там есть
	check(err)
	err = file.Close()
	check(err)
	fmt.Println("Файл обновлён!")
}
