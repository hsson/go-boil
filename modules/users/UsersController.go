package users

import (
  "net/http"
  "fmt"
)

// POST: /users/sign_in
func SignIn(w http.ResponseWriter, r *http.Request) {
  username := r.FormValue("username")
  password := r.FormValue("password")
  user, err := findUserByUsername(username, getDummyData())

  if (err != nil) {
    http.Error(w, "User not found", http.StatusUnauthorized)
    return
  }

  // TODO: Hash passwords
  if (user.Password != password) {
    http.Error(w, "Password is not correct", http.StatusUnauthorized)
    return
  }

  fmt.Fprintf(w, "Successfully signed in %s\n", user.Name)
}
