package main

import "fmt"

func main() {
	array := []int{0}
	fmt.Println(array)
	for i := 1; i < 10; i++ {
		array = append(array, i)
	}
	fmt.Println(array)
}
