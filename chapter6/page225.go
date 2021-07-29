package main

//bcxyz
import (
	"fmt"
)

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	slice := array[1:3]
	slice = append(slice, "x")
	slice = append(slice, "y", "z")
	for _, letter := range slice {
		fmt.Println(letter)
	}
}
