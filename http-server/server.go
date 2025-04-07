package main

import (
	"fmt"
	"net/http"
	"strings"
)

type InMemoryStore struct{}

func (i *InMemoryStore) GetPlayerScore(playerName string) int {
	return 123
}

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}
