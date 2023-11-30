package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jon")
	want := "Olá, Jon"

	if got != want {
		t.Errorf("git %q want %q", got, want)
	}
}
