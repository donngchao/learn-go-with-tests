package main

import (
	"fmt"
	"net/http"
)

// PlayerServer currently returns Hello, world given _any_ request
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "40")
}

func MiniServer(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprint(w,"simple output")
}

