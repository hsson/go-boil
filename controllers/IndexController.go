package controllers

import (
  "fmt"
  "net/http"
)

// GET: /
func GetIndex(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello index!")
}
