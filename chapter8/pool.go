package main

import (
	"fmt"
)

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

func main() {
	location := Coordinates{37.42, -122.08}
	fmt.Println("Lat:", location.Latitude, "\tLong:", location.Longitude)
}
