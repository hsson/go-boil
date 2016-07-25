package main

import (
  "github.com/urfave/negroni"
)

func main() {
	r := NewRouter()
	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":8080")
}
