package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type InMemoryStore struct {
	scores map[string]int
}

func (s *InMemoryStore) GetPlayerScore(playerName string) (int, error) {
	if score, ok := s.scores[playerName]; ok {
		return score, nil
	} else {
		err := errors.New("player doesn't exist")
		return 0, err
	}
}

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score, err := p.store.GetPlayerScore(player)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Fprint(w, score)
	}
}
