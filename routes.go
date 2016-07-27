package main

import (
	"net/http"

	"github.com/hsson/go-boil/modules/messages"
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
	Route{"Index", "GET", "/", false, IndexSite},

	// Messages
	Route{"Messages", "GET", "/messages/", false, messages.Index},
	Route{"Messages", "GET", "/messages/{id:[0-9]+}", false, messages.Show},
}
