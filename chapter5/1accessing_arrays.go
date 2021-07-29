/*
Exercise 1: Accessing Arrays

Create an array with 5 string elements, holding English weekday names:
“Monday”, “Tuesday”, “Wednesday”, “Thursday”, “Friday”.
Then print each array element along with its index.

You can assign array elements individually, or you can use an array literal.
You can access the elements individually, use a for loop to get each element index,
or use a for ... range loop to loop over the elements themselves.
Better yet, try all these techniques!
*/
package main

import (
	"fmt"
)

func main() {
	var week [5]string = [5]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	for i, elem := range week {
		fmt.Println(i, elem)
	}

	//OR
	var nedelya [5]string = [5]string{} //"ПН", "ВТ", "СР", "ЧТ", "ПТ"
	nedelya[0] = "ПН"
	nedelya[1] = "ВТ"
	nedelya[2] = "СР"
	nedelya[3] = "ЧТ"
	nedelya[4] = "ПТ"
	for j := 0; j < len(nedelya); j++ {
		fmt.Println(j, nedelya[j])
	}
}
