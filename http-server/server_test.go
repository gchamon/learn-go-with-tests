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

func TestGETPlayers(t *testing.T) {
	server := &PlayerServer{}

	t.Run("returns Pepper's score", func(t *testing.T) {
		assertPlayerScore(t, server, "Pepper", "20")
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		assertPlayerScore(t, server, "Floyd", "10")
	})
}
