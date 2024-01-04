package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Dane")
	want := "Hello, Dane"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}