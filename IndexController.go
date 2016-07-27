package main

import (
	"fmt"
	"net/http"
)

// GET: /
func IndexSite(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello index!")
}
