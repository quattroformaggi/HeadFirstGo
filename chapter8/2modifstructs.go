/*Exercise 2: Modifying Structs from Functions

All squares are rectangles, but not all rectangles are squares…
Let’s define a makeSquare function that takes a rectangle
and “cuts it down” so that its longer sides are equal to its shorter sides.

If the rectangle’s length is greater than its width, set its length equal to its width.
Otherwise, set the width equal to the length.

makeSquare won’t return a value; it should modify the rectangle it receives
(meaning it will need to accept a pointer to a rectangle and modify the value at that pointer).

In main, create a couple different rectangle values,
one where the length is greater and one where the width is greater, and try converting them to squares using makeSquare.*/
package main

import "fmt"

type Rectangle struct {
	Length float64
	Width  float64
}

func rekt(rek Rectangle) {
	fmt.Println(rek.Length, "x", rek.Width)
}

func makeSquare(rptr *Rectangle) {
	if rptr.Length > rptr.Width {
		rptr.Length = rptr.Width
	} else {
		rptr.Width = rptr.Length
	}
}

func main() {
	var r Rectangle
	r.Length = 10
	r.Width = 11
	makeSquare(&r)
	rekt(r)
	var re Rectangle
	re.Length = 9
	re.Width = 11
	makeSquare(&re)
	rekt(re)
}
