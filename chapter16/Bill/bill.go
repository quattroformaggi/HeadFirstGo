package main

import (
	"html/template"
	"os"
	"log"
)

type Invoice struct{
	Name string
	Paid bool
	Charges []float64
	Total float64
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main(){
	html, err := template.ParseFiles("bill.html")
	check(err)
	bill:=Invoice{
		Name: "Мэри Сью",
		Paid: true,
		Charges: []float64{25.0,1.5,44.44},
		Total: 70.94,
	}
	err = html.Execute(os.Stdout,bill)
	check(err)
}