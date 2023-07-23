package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Enrique")
	want := "Hello, Enrique"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
