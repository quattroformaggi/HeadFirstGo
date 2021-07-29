package main

import (
	"html/template"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func writeList(writer http.ResponseWriter, list []string) {
	tmp, err := template.ParseFiles("list.html")
	check(err)
	err = tmp.Execute(writer,list)
	check(err)
	// YOUR CODE HERE: Use the template library to parse the contents
	// of the "list.html" file. You'll get a template value and an
	// error value.
	// Pass the error value to the "check" function.
	// Now call the "Execute" method on the template. It should write
	// its output to the "writer" parameter, and it should use the
	// "list" parameter as its data value.
	// You'll get another error value back from "Execute", which
	// should be passed to "check".
}

func fruitHandler(writer http.ResponseWriter, request *http.Request) {
	writeList(writer, []string{"яблоки", "апельсины", "груши"})
}

func meatHandler(writer http.ResponseWriter, request *http.Request) {
	writeList(writer, []string{"курица", "говядина", "баранина"})
}

func main() {
	http.HandleFunc("/fruit", fruitHandler)
	http.HandleFunc("/meat", meatHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}