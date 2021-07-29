/*Several of the variable names used in this example break Go naming conventions.
The code is also somewhat longer than it needs to be.
See if you can modify so that it’s shorter
and follows conventions better,
but still produces the same output.*/
package main

import "fmt"

func main() {
	var pebbleweight, rockweight, boulderweight float64 = 0.1, 1.2, 502.4
	var total_weight = pebbleweight + rockweight + boulderweight
	fmt.Println(total_weight)
}
