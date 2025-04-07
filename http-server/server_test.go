package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func postPlayerScore(server *PlayerServer, playerName string, score int) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(
		http.MethodPost,
		"/players/"+playerName,
		nil,
	)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	return response
}

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

func TestGETPlayers(t *testing.T) {
	store := InMemoryStore{
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
			t.Errorf("got status %d want %d", got, want)
		}
	})
}

func TestStoreWins(t *testing.T) {
	store := InMemoryStore{
		map[string]int{},
	}
	server := &PlayerServer{&store}

	t.Run("returns accepted on POST", func(t *testing.T) {
		response := postPlayerScore(server, "Pepper", 20)

		got := response.Code
		want := http.StatusAccepted

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})
}
