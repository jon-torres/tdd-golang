package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jon")
	want := "Hello, Jon"

	if got != want {
		t.Errorf("git %q want %q", got, want)
	}
}
