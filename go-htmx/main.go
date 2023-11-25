package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Film struct {
	Title string
	Director string
}


func main()  {
	h1 := func (w http.ResponseWriter, r *http.Request)  {
		// io.WriteString(w, "Hello world \n")
		// io.WriteString(w, r.Method)
		temp1 := template.Must(template.ParseFiles("index.html"))

		// Pass context to the template, nil for no context to be passed
		temp1.Execute(w, nil)
	}

	http.HandleFunc("/", h1)

	fmt.Println("Serving on port 6060...")
	log.Fatal(http.ListenAndServe(":6060", nil))
}