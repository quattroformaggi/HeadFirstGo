/*Exercise 1: Defined Types and Structs
Let’s create a struct type named rectangle, that represents rectangular areas. It should have length and width fields, both of type float64.
Also create a rectangleInfo function that accepts a rectangle as a parameter.
	rectangleInfo should print "Length:" followed by the rectangle’s length, then "Width:" followed by the rectangle’s width.
Finally, in your main function, create a new rectangle value, and set its length and width fields. Then pass the rectangle to rectangleInfo to display its field values.*/
package main

import (
	"fmt"
)

type Rectangle struct {
	Length float64
	Width  float64
}

func rectangleInfo(rec Rectangle) {
	fmt.Println("Длина:", rec.Length, "\nШирина:", rec.Width)
}

func main() {
	var r Rectangle
	r.Length = 10
	r.Width = 11
	rectangleInfo(r)
}
