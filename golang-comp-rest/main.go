
package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Movie struct {
	Title string `json:"title"`
	Timelimit int `json:"timelimit"`
	Cast string `json:"cast"`
}

var Movies []Movie

func Intro(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Let's start")
	fmt.Println("Endpoint Hit: Movie Details")
}

func returnAllMovies(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: returnAllMoviews")
	json.NewEncoder(w).Encode(Movies)
}

func handleRquests()  {
	http.HandleFunc("/", Intro)
	http.HandleFunc("/movies", returnAllMovies)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func main()  {
	Movies = []Movie {
		Movie{Title: "Rambo", Timelimit: 1234, Cast: "Sylvester stallon"},
		Movie{Title: "Bridge of spies", Timelimit: 3443, Cast: "John wick"},
	}
	handleRquests()
}