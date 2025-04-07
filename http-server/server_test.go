package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getPlayerScore(server *PlayerServer, playerName string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(
		http.MethodGet,
		"/players/"+playerName,
		nil,
	)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	return response
}

func assertPlayerScore(t testing.TB, server *PlayerServer, playerName string, want string) {
	response := getPlayerScore(server, playerName)

	gotStatus := response.Code
	wantStatus := http.StatusOK
	got := response.Body.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if gotStatus != wantStatus {
		t.Errorf("got status %d, want %d", gotStatus, wantStatus)
	}
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(playerName string) (int, error) {
	if score, ok := s.scores[playerName]; ok {
		return score, nil
	} else {
		err := errors.New("player doesn't exist")
		return 0, err
	}
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

	t.Run("returns 404 on missing players", func(t *testing.T) {
		response := getPlayerScore(server, "Apollo")

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
