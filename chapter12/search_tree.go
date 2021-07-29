package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func scanDir(path string) error {
	fmt.Println("Путь:", path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDir(filePath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

func main() {
	err := scanDir("mydir")
	if err != nil {
		log.Fatal(err)
	}
}
