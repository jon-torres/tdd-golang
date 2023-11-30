package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying 'olá' to people", func(t *testing.T) {
		got := Hello("Jon")
		want := "Olá, Jon"

		if got != want {
			t.Errorf("git %q want %q", got, want)
		}
	})
	t.Run("Say 'Olá, Mundo' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Olá, Mundo"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
