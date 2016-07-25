package main

import (
  "net/http"
  "github.com/urfave/negroni"
)

func main() {
	r := NewRouter()
	n := negroni.New()

  // Recovery handling middleware
  n.Use(negroni.NewRecovery())

  // Logging middleware
  n.Use(negroni.NewLogger())

  // Static file serving middleware
  n.Use(negroni.NewStatic(http.Dir("public")))

  n.UseHandler(r)
	n.Run(":8080")
}
