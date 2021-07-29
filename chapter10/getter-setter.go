package main

import (
	"calendar"
	"fmt"
	"log"
)

func main() {
	d := calendar.Date{}
	err := d.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}
	err = d.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = d.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("D:", d.Day(), "M:", d.Month(), "Y:", d.Year())
}
