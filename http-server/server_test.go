package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertPlayerScore(t testing.TB, server *PlayerServer, playerName string, want string) {
	request, _ := http.NewRequest(
		http.MethodGet,
		"/players/"+playerName,
		nil,
	)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	got := response.Body.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(playerName string) int {
	score := s.scores[playerName]
	return score
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		assertPlayerScore(t, server, "Pepper", "20")
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		assertPlayerScore(t, server, "Floyd", "10")
	})
}
