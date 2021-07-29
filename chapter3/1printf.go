package main

import "fmt"

func scoreSummary(s string, pts1 float64, pts2 float64, pts3 float64) {
	var avg float64 = (pts1 + pts2 + pts3) / 3
	fmt.Printf("%10s\t| %8.2f\t| %8.2f\t| %8.2f\t| %8.2f", s, pts1, pts2, pts3, avg)
	fmt.Println()
}

func main() {
	fmt.Printf("%10s | %8s | %8s | %8s | %8s\n",
		"Имя", "Русский", "Математика", "ИКТ", "Average")
	for i := 0; i < 35; i++ {
		fmt.Print("-")
	}
	fmt.Println("------------------------------------------------------")
	scoreSummary("Сергей", 95.4, 82.3, 74.6)
	scoreSummary("Кирилл", 79.3, 99.1, 82.5)
	scoreSummary("Антон", 82.2, 95.4, 77.6)
}

//% ширина-числа.ширина-др-части тип-глагола-формат

/*Exercise 1: The program below should produce a table with student names, their scores on three tests, and the average of those scores.
Write a scoreSummary function that accepts a string parameter with the student name, followed by three float64 parameters with the test scores.
Calculate the average of the scores by adding them all together and dividing the result by three.
Then call fmt.Printf to print a table row with the following columns:
    The student name string, with a width of 10 characters.
    The first test score, with a with of 8 characters, 2 of them decimal places.
    The second test score, with a with of 8 characters, 2 of them decimal places.
    The third test score, with a with of 8 characters, 2 of them decimal places.
    The average of the scores, with a with of 8 characters, 2 of them decimal places.

Separate the columns with a vertical bar surrounded by spaces, just like you see in the call to Printf in the main function.*/
