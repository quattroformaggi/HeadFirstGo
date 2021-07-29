/*Exercise 2: Query Parameters

We’ve set up a getParameter function for you, which can read the value of a query parameter.

Your task is to use getParameter in a web app.
You’ll be writing a request handler function that takes a query parameter and displays it as an HTML <h1> heading.

Set up a handler function that can be passed to http.HandleFunc (that is, it must accept http.ResponseWriter and *http.Request parameters).
    Within the function, call getParameter to get the value of the “text” parameter.
    Write the returned string out the the response in an <h1> HTML tag.
Within main:
    Set up your function to handle requests for the “/big” path.
    Then start an HTTP server on port 8080.*/
package main

import (
	"log"
	"net/http"
)

// getParameter returns the first value associated with a
// named query parameter from an http.Request.
func getParameter(request *http.Request, parameterName string) string {
	query := request.URL.Query()
	return query[parameterName][0]
}

// YOUR CODE HERE: Set up a handler function.
func paramHandler(writer http.ResponseWriter, request *http.Request) {
	text := getParameter(request, "TEXT")
	body := []byte("<h1>" + text + "</h1>")

	_, err := writer.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// YOUR CODE HERE: Set up your function to handle
	// requests for the "/big" path.
	// Then start an HTTP server on port 8080.
	http.HandleFunc("/big", paramHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
