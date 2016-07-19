package main

import (
  "net/http"
  "fmt"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()

	n := negroni.Classic()

  r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
  })
	n.UseHandler(r)
	n.Run(":8080")
}
