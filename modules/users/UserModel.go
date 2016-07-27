package users

type User struct {
  ID        int       `json:"id"`
  Username  string    `json:"usernames"`
  Password  string    `json:"-"`
  Name      string    `json:"name"`
}
