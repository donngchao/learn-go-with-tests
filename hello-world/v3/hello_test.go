package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHi(t *testing.T) {
	got := Hi("Mike")
	want := "Hi, Mike"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
