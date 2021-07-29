package main

import "fmt"

func main() {
	lyrics := [4]string{"да,", "я", "богатый", "лаваш"}
	slice1 := lyrics[1:3]
	fmt.Println(slice1) //я богатый
	slice2 := lyrics[1:]
	fmt.Println(slice2) //я богатый лаваш
	slice3 := lyrics[:2]
	fmt.Println(slice3) //да, я
}
