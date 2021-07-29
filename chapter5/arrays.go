//Package greetings has custom greet functions (for chapter #4)
package main

import (
	"fmt"
)

func main() {
	var words [4]string = [4]string{"Жижа\n", "в", "столовке", "meh..."}
	fmt.Println(words)
	fmt.Printf("%#v\n", words)
}
