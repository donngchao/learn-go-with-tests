package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore stores score information about players
type PlayerStore interface {
	//GetPlayerScore is a function that takes in string returns int
	GetPlayerScore(name string) int
	PrintOutScore(name string)
}

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//去掉访问路径前缀的/players/
	player := strings.TrimPrefix(r.URL.Path, "/players/")

    fmt.Println(r.URL.Path)
	fmt.Println(player)
	score := p.store.GetPlayerScore(player)
	p.store.PrintOutScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
