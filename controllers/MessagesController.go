package controllers

import (
  "github.com/hsson/go-appengine-boilerplate/models"

  "fmt"
  "net/http"
  "encoding/json"
)

func IndexMessages(w http.ResponseWriter, r *http.Request) {
  message := models.Message{"hsson", "This is an example."}
  response, err := json.MarshalIndent(message, "", "\t")
  if (err != nil) {
    http.Error(w, "Something went wrong, try again.", http.StatusInternalServerError)
    return
  }
  fmt.Fprint(w, string(response))
}
