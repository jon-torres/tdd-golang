package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say 'Olá, Mundo' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Olá, Mundo"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying 'olá' to a person", func(t *testing.T) {
		got := Hello("Jon", "")
		want := "Olá, Jon"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Pierre", french)
		want := "Bonjour, Pierre"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Maria", spanish)
		want := "Hola, Maria"
		assertCorrectMessage(t, got, want)
	})
}
func assertCorrectMessage(t testing.TB, got, want string) {
	// if a test fails, by adding the Helper() method, it points
	// to the error inside the "main" function, making it easier to debug
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
