package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHelloServer(t *testing.T) {
  host, _ := os.Hostname()

  request, _ := http.NewRequest(http.MethodGet, "/", nil)
  response := httptest.NewRecorder()

  HelloServer(response, request) 

  got := response.Body.String()
  want := "Hello, " + host

  if got != want {
    t.Errorf("got %q, want %q", got, want)
  }
}
