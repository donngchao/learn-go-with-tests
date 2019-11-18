package main

import (
	"log"
	"net/http"
)

func main() {

/*	type PlayerServer struct {
		store PlayerStore
	}
*/

/**
  type PlayerStore interface {
  	GetPlayerScore(name string) int
  	RecordWin(name string)
  }
*/
	server := &PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
