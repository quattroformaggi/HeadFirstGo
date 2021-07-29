// one project main.go
package main

import (
	"fmt"
	"math"
)

func main() {
	var one float64
	var two float64
	var choice int
	fmt.Println("Enter two numerals with a space in between:")
	fmt.Scan(&one, &two)
	fmt.Println("1. +\t2. -\t3. *\t4. /\t5. ^\t6. V\t7. %\n\tEnter the number of desired operation:")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("Sum:\t", one, "+", two, "=", one+two)
		break
	case 2:
		fmt.Println("Difference:\t", one, "-", two, "=", one-two)
		break
	case 3:
		fmt.Println("Product:\t", one, "*", two, "=", one*two)
		break
	case 4:
		fmt.Println("Quotient:\t", one, "/", two, "=", one/two)
		break
	case 5:
		fmt.Println("Power:\t", one, "^", two, "=", math.Pow(one, two))
		break
	case 6:
		fmt.Println("Exponentiation:\tV", one, "=", math.Sqrt(one), "\tV", two, "=", math.Sqrt(two))
		break
	case 7:
		fmt.Println("Percentage:\t", one, "%", two, "=", math.Ceil(one/two*100))
		break
	default:
		fmt.Println("ERROR: Undefined operation code.")
		break
	}
	fmt.Println("Program finished successfully.")
}
