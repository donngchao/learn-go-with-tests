package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	PlayerServer(response, request)

	t.Run("returns Pepper's score", func(t *testing.T) {
		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}

func TestMiniServer(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	MiniServer(response, request)

	t.Run("returns simple output", func(t *testing.T) {
		got := response.Body.String()
		want := "simple output"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
