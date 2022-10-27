package main

import (
	"bytes"
	"testing"
)

func TestGreat(t *testing.T) {
	t.Run("Great Test", func(t *testing.T) {
		buffer := bytes.Buffer{}

		Greet(&buffer, "Sachin")

		got := buffer.String()
		want := "Hello, Sachin"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
