package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	// slowServer := makeDelayServer(20 * time.Millisecond)
	// fastServer := makeDelayServer(0 * time.Millisecond)

	// defer slowServer.Close()
	// defer fastServer.Close()

	// slowURL := slowServer.URL
	// fastURL := fastServer.URL

	// want := fastURL
	// got, _ := Racer(slowURL, fastURL)

	// if got != want {
	// 	t.Errorf("got %q, want %q", got, want)
	// }

	// slowServer.Close()
	// fastServer.Close()

	t.Run("return an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayServer(11 * time.Second)
		serverB := makeDelayServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err != nil {
			t.Error("expected an error but didn't get one", err.Error())
		}
	})

}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
