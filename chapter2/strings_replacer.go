package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "Go is * function*l l*ngu*ge."
	fmt.Println("Old string:", broken)
	replacer := strings.NewReplacer("*", "a")
	fixed := replacer.Replace(broken)
	fmt.Println("New string:", fixed)
}
