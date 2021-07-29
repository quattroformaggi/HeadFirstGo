package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func scanDir(path string) error {
	fmt.Println("Путь:", path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDir(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

func main() {
	scanDir("mydir")	//no need to check "err" val
}
