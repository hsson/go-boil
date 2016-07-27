package users

import (
  "errors"
)

func getDummyData() []User {
  return []User{
    User{
      ID: 1,
      Username: "john",
      Password: "password123",
      Name: "John Doe",
    },
  }
}

func findUserByUsername(username string, users []User) (User, error) {
  for _,u := range users {
    if (u.Username == username) {
      return u, nil
    }
  }
  return User{}, errors.New("No user found")
}
