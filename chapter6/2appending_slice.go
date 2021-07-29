package main

/*Exercise 2: Appending to a Slice
The program below should ask the user to input several strings.
When the user is done, it should output the strings the user
entered, separated by commas.*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//This allows us to set up a reader just once,
// at the start of the "main" function, and then keep
// re-using it for the life of the program.
var reader *bufio.Reader

// getString prompts the user to input a string. Returns
// an empty string and an error if there was a problem
// getting input, or the user's input and a nil error
// otherwise.
func getString(question string) (string, error) {
	// Display the question and prompt for the answer
	// on the same line.
	fmt.Print(question + ": ")
	input, err := reader.ReadString('\n')
	// If there was an error, just return it to the caller.
	if err != nil {
		return "", err
	}
	// Remove the newline from the end.
	input = strings.TrimSpace(input)
	// Return the input and a "nil" error.
	return input, nil
}

// askToContinue prompts the user if they want to continue.
// If they enter "y", "Y", or no input, it will return true.
// If they enter "n" or "N", it will return false.
func askToContinue() bool {
	// This uses a "for" loop without any condition, which
	// just loops forever (until the user enters an accepted
	// answer and we return).
	for {
		answer, err := getString("Continue? [Y/n]")
		// If there was an error, log it and exit.
		if err != nil {
			log.Fatal(err)
		}
		// If user entered "Y" or "N", convert it to lower-case.
		answer = strings.ToLower(answer)
		// If user entered "n", don't continue.
		if answer == "n" {
			return false
			// If user entered "y" or nothing, continue.
		} else if answer == "y" || answer == "" {
			return true
		}
	}
}

func main() {
	// Store the Reader that getString will use in the
	// package variable. Note this is an assignment, not a
	// declaration!
	reader = bufio.NewReader(os.Stdin)
	// Process at least one phrase
	more := true
	// YOUR CODE HERE: Set up a slice to hold phrases the
	// user enters.
	var phrases []string

	// Loop until user chooses not to continue.
	for more == true {
		// Get the user's input.
		input, err := getString("String to add")
		// If there was an error, log it and exit.
		if err != nil {
			log.Fatal(err)
		}
		phrases = append(phrases, input)

		// Ask if the user wants to continue.
		more = askToContinue()
	}
	strings.Join(phrases, ", ")
	fmt.Println("\nFinalized String:", phrases)
}
