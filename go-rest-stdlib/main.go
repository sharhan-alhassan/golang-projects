
package main

import "net/http"

type homeHandler struct {}
type RecipeHandler struct {}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write( []byte ("This is my Homepage"))
}

func (h *RecipeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my recipe page"))
}
func main()  {
	mux := http.NewServeMux()

	mux.Handle("/", &homeHandler{})
	mux.Handle("/recipes", &RecipeHandler{})
}

