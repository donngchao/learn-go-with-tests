package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestNihao(t *testing.T) {
	got := Nihao("GO")
	want := "你好, GO"

	if got != want {
		t.Errorf("got %q want %q",got,want)
	}
}
