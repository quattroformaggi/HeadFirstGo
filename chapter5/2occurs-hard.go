/*Exercise 2: Counting Occurrences

We’ve set up a small array that contains the numbers 0, 1, and 2,
each repeated a random number of times.
Update the program to count how many times each number occurs.*/
package main

import "fmt"

func main() {
	// We'll count the number of times each number occurs
	// within this array.
	numbers := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	// "occurrences" will have the indexes 0, 1, and 2.
	var occurrences [3]int
	// Process each element in "numbers". We ignore the
	// index because we don't need it.
	for _, number := range numbers {
		// The zero value for elements in "occurrences"
		// is 0. So we can safely increment it even if
		// we've never assigned to it.
		occurrences[number]++
	}
	// Finally, print out the number of times each number
	// occurred.
	for number, count := range occurrences {
		fmt.Println(number, "occurred", count, "times")
	}
	fmt.Println()
	//////////////////////////////////////////////
	//////////////////////////////////////////////
	//////////////////////////////////////////////

	nums := [100]int{4, 2, 6, 7, 8, 0, 1, 8, 7, 8,
		1, 5, 3, 2, 2, 1, 9, 6, 1, 0, 0, 0, 8, 4, 6,
		2, 2, 4, 7, 9, 6, 5, 9, 0, 5, 1, 1, 5, 4, 7,
		7, 9, 7, 8, 6, 3, 3, 3, 4, 8, 0, 4, 1, 1, 7,
		9, 8, 8, 1, 2, 3, 6, 4, 9, 2, 5, 8, 6, 7, 7,
		5, 4, 2, 9, 4, 4, 2, 2, 5, 5, 0, 0, 0, 9, 1,
		9, 5, 8, 0, 1, 1, 0, 5, 3, 8, 6, 3, 4, 4, 9}
	var occurs [10]int
	for _, num := range nums {
		occurs[num]++
	}
	for num, coun := range occurs {
		fmt.Println(num, "occurred", coun, "times")
	}
}
