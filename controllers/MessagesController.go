package controllers

import (
  "github.com/hsson/go-appengine-boilerplate/models"
  "github.com/gorilla/mux"

  "fmt"
  "net/http"
  "encoding/json"
)

// GET: /messages/
func IndexMessages(w http.ResponseWriter, r *http.Request) {
  message := models.Message{"hsson", "This is an example."}
  response, err := json.MarshalIndent(message, "", "\t")
  if (err != nil) {
    http.Error(w, "Something went wrong, try again.", http.StatusInternalServerError)
    return
  }
  fmt.Fprint(w, string(response))
}

// GET: /messages/1
func ShowMessage(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  fmt.Fprintf(w, "The id is: %s", id)
}
