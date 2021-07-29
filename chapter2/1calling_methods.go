/*Exercise 1: Calling Methods

The strings package has a Builder type that is used to build a long text string out of multiple shorter strings. A strings.Builder value has a variety of different methods you can call, but this exercise is going to focus on these three:

    The WriteRune method accepts a rune as an argument and adds that rune onto the end of the string being built.
    The WriteString method works just like WriteRune, except that it accepts a string as an argument.
    The String method takes no arguments. Its return value is the accumulated string.

Complete the code sample below. The sample contains YOUR CODE HERE comments in the places you should add code, with a description of what your code should do.*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	// The "builder" variable holds a "strings.Builder"
	// value. We're not assigning a value to "builder",
	// but that's okay; you can call methods on the
	// zero value for the "strings.Builder" type.
	var builder strings.Builder

	// Call the "WriteRune" method on
	// "builder" 3 times, once with the rune 'a', once
	// with the rune 'b', and a third time with 'c'.
	builder.WriteRune('a')
	builder.WriteRune('b')
	builder.WriteRune('c')

	// Call the "WriteString" method on
	// "builder" once, with the string "defg".
	builder.WriteString("defg")

	// Call the "String" method on
	// "builder", and pass the return value to
	// "fmt.Println".
	fmt.Println(builder.String())
}
