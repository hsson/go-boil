package main

import (
  "net/http"
)

type Route struct {
  Name        string
  Method      string
  Pattern     string
  Secured     bool
  HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
  // Index page
  Route{"Index", "GET", "/", false, GetIndex},
}
