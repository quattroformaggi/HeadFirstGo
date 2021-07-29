package main

import (
	"fmt"
)

func main() {
	jewelry := make(map[string]float64)
	jewelry["necklace"] = 89.99
	jewelry["earrings"] = 79.99
	clothing := map[string]float64{"pants": 59.99, "shirt": 39.99}
	fmt.Println("Серёжки:", jewelry["earrings"])
	fmt.Println("Ожерелье:", jewelry["necklace"])
	fmt.Println("Майка:", clothing["shirt"])
	fmt.Println("Штаны:", clothing["pants"])
}
