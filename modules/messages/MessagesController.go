package messages

import (
  "github.com/gorilla/mux"

  "strconv"
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
  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

// GET: /messages/1
func ShowMessage(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])

  // This shouldn't happen since the mux only accepts numbers to this route
  if (err != nil) {
    http.Error(w, "Invalid Id, please try again.", http.StatusBadRequest)
    return
  }

  messages := getDummyData()
  var message Message
  found := false

  // Super simple but crap search
  for _,m := range messages {
    if (m.Id == id) {
      message = m
      found = true
      break
    }
  }

  if (!found) {
    http.NotFound(w, r)
    return
  }

  response, err := json.MarshalIndent(message, "", "\t")
  if (err != nil) {
    http.Error(w, "Something went wrong, try again.", http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

func getDummyData() []Message {
  data := make([]Message, 10)
  for i := 0; i < 10; i++ {
    data[i] = Message{Id: i + 1, Author: "Alexander", Message: "Random message " + strconv.Itoa(i + 1)}
  }
  return data
}
