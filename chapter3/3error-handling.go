package main

import (
	"fmt"
	"log"
)

func perimeter(wid float64, hei float64) (float64, error) {
	if wid < 0 {
		return 0, fmt.Errorf("Отрицательная длина невозможна!:", wid)
	}
	if hei < 0 {
		return 0, fmt.Errorf("Отрицательная высота невозможна!:", hei)
	}

	return wid * hei * 2, nil
}

func main() {
	// YOUR CODE HERE: Call "perimeter" three times.
	// Add the three return values together, and store the
	// total in the "total" variable.
	total, err := perimeter(-8.2, 10.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Вам потребуется", total, "метров заграждения.")
}

/*
Exercise 3: Error Handling

Update your perimeter function from the previous exercise to return an error if either the length or width are negative.

    If the given length is negative, return a perimeter of 0 and an error with the message “a length of [length] is invalid”.
    If the given width is negative, return a perimeter of 0 and an error with the message “a width of [width] is invalid”.
    Otherwise, return the calculated perimeter, and an error value of nil.

In your main function, test returning an error by calling perimeter with a negative value. If the returned error value is not nil, print the error message and exit the program. Otherwise, print the returned perimeter.

Sample output:*/
