package users

import (
  jwt "github.com/dgrijalva/jwt-go"
  "net/http"
  "fmt"
  "time"
)

// A secret key for signing tokens
const tokenSecretKey string = "super-secret-key"

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

  // Sign in is successfull, generate JWT
  token := jwt.New(jwt.SigningMethodHS256)
  token.Claims = map[string]interface{} {
    "username": user.Username,
    "name": user.Name,
    "iat": time.Now().Unix(),
    "exp": time.Now().Add(24 * 30 * time.Hour).Unix(),
  }

  tokenString, err := token.SignedString([]byte(tokenSecretKey))
  if (err != nil) {
    fmt.Fprintln(w, err)
    http.Error(w, "Something went wrong, try again", http.StatusInternalServerError)
    return
  }

  fmt.Fprint(w, tokenString)
}
