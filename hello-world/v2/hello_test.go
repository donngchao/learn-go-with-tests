package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHi(t *testing.T) {
	got := Hi()
	want := "Hi,there! I am using golang"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

