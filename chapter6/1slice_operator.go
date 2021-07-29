package main

import "fmt"

/*Exercise 1: The Slice Operator
Define a printLines function that takes a slice of strings as a parameter,
and prints each element of that slice on a separate line.
Then, in main, get a slice of the daysOfWeek array containing just the weekdays:
“Monday”, “Tuesday”, “Wednesday”, “Thursday”, and “Friday”.
Pass that slice to printLines.*/

func printLines(strings []string) {
	for _, a := range strings {
		fmt.Println(a)
	}
}

func main() {
	daysOfWeek := [7]string{"Sunday", "Monday", "Tuesday",
		"Wednesday", "Thursday", "Friday", "Saturday"}
	five := daysOfWeek[1:6]
	printLines(five)
}
