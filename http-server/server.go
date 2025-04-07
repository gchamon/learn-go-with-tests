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

func (p *PlayerServer) processWin(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score, err := p.store.GetPlayerScore(player)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Fprint(w, score)
	}
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.processWin(w)
	case http.MethodGet:
		p.showScore(w, r)
	}
}
