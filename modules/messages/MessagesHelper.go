package messages

import (
  "strconv"
  "net/http"
  "encoding/json"
  "errors"
)

func getDummyData() []Message {
  data := make([]Message, 10)
  for i := 0; i < 10; i++ {
    data[i] = Message{ID: i + 1, Author: "Alexander", Message: "Random message " + strconv.Itoa(i + 1)}
  }
  return data
}

func printJSON(message interface{}, w http.ResponseWriter) {
  response, err := json.MarshalIndent(message, "", "\t")
  if (err != nil) {
    http.Error(w, "Something went wrong, try again.", http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

func getMessagesByID(messages []Message, id int) (Message, error) {
  var message Message
  found := false
  // Super simple but crap search
  for _,m := range messages {
    if (m.ID == id) {
      message = m
      found = true
      break
    }
  }

  if (!found) {
    return Message{}, errors.New("Message not found")
  }
  return message, nil
}
