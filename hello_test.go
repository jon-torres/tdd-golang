package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying 'olá' to people", func(t *testing.T) {
		got := Hello("Jon")
		want := "Olá, Jon"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Olá, Mundo' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Olá, Mundo"
		assertCorrectMessage(t, got, want)
	})
}
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
