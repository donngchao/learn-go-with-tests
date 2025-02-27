package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(got, want)
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", spanish)
		want := "Hola, Elodie"
		assertCorrectMessage(got, want)
	})

	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Lauren", french)
		want := "Bonjour, Lauren"
		assertCorrectMessage(got, want)
	})

	t.Run("say hello in chinese", func(t *testing.T) {
		got := Hello("张三", china)
		want := "你好, 张三"
		assertCorrectMessage(got, want)
	})

}
