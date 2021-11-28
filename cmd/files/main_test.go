package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("200 ok", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/Go", nil)
		response := httptest.NewRecorder()
		Handler(response, request)
		got := response.Body.String()
		want := "Hi there, I love Go!"
		if got != want {
				t.Errorf("got %q, want %q", got, want)
		}
	})
}
