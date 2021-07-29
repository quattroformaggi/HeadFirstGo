/*Exercise 1: Counting Occurrences of Strings
In the chapter 5 exercises, we asked you to count how many times the numbers 0, 1, and 2 occurred within an array.
This time, your task is to count how many times letters occur within a slice.*/
package main

import "fmt"

func main() {
	// We'll count the number of times each letter occurs
	// within this slice.
	letters := []string{"b", "a", "c", "a", "b", "a",
		"a", "b", "c"}
	var counters [3]int
	for _, elem := range letters {
		switch elem {
		case "a":
			counters[0] += 1
			break
		case "b":
			counters[1] += 1
			break
		case "c":
			counters[2] += 1
			break
		}
	}
	fmt.Println("A:B:C occurred", counters, "times")
	// YOUR CODE HERE: Process each element of "letters",
	// keeping track of how many times you've seen "a",
	// "b", or "c".
	// Finally, print out the number of times each letter
	// occurred.
}
