package messages

import (
  "github.com/gorilla/mux"

  "strconv"
  "net/http"
)

// GET: /messages/
func Index(w http.ResponseWriter, r *http.Request) {
  printJSON(getDummyData(), w)
}

// GET: /messages/1
func Show(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])

  // This shouldn't happen since the mux only accepts numbers to this route
  if (err != nil) {
    http.Error(w, "Invalid Id, please try again.", http.StatusBadRequest)
    return
  }

  messages := getDummyData()
  message, err := getMessagesByID(messages, id)

  if (err != nil) {
    http.NotFound(w, r)
    return
  }

  printJSON(message, w)
}
