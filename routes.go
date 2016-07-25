package main

import (
  "net/http"

  "github.com/hsson/go-appengine-boilerplate/controllers"
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
  Route{"Index", "GET", "/", false, controllers.GetIndex},

  // Messages
  Route{"Messages", "GET", "/messages/", false, controllers.IndexMessages},
}
