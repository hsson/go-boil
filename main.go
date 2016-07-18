package main

import (
  "net/http"

  "github.com/urfave/negroni"
  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()

  n := negroni.Classic()
  n.UserHandler(r)
  n.Run(":8080")
}
