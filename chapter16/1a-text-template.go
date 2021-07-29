/*Exercise 1: A Text Template

We want to create a text template that will output a bulleted list
(using simple dash characters as the bullets).
Set the templateText variable so that the program will produce the output shown.*/
package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func execTemplate(templateText string, data interface{}) {
	tmpl, err := template.New("template").Parse(templateText)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	templateText := "{{range .}}->\t{{.}}\n{{end}}"
	execTemplate(templateText, []string{"яблоки", "апельсины", "груши"})
	execTemplate(templateText, []string{"курица", "говядина", "телятина"})
}
