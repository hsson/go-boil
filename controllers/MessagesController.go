package controllers

import (
  "github.com/hsson/go-appengine-boilerplate/models"
  "github.com/gorilla/mux"

  "strconv"
  "fmt"
  "net/http"
  "encoding/json"
)

// GET: /messages/
func IndexMessages(w http.ResponseWriter, r *http.Request) {
  response, err := json.MarshalIndent(getDummyData(), "", "\t")
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

func getDummyData() []models.Message {
  data := make([]models.Message, 10)
  for i := 0; i < 10; i++ {
    data[i] = models.Message{Id: i + 1, Author: "Alexander", Message: "Random message " + strconv.Itoa(i + 1)}
  }
  return data
}
