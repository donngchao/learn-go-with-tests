package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	apple := Dictionary{"apple":"a kind of fruit"}
	p := PhoneBook{"Jack":130110}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("check apple", func(t *testing.T) {
		got, _ := apple.Search("apple")
		want := "a kind of fruit"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})

	t.Run("another unknown word", func(t *testing.T) {
		_, got := dictionary.Search("cheek")

		assertError(t, got, ErrNotFound)
	})

	t.Run("find Jack in the PhoneBook", func(t *testing.T) {
		got, _ := p.Search("Jack")
		want := 130110

		assertStrings2(t, got, want)
	})

	t.Run("find Gimmy in the PhoneBook", func(t *testing.T) {
		got, _ := p.Search("Gimmy")
		assertError2(t, got, 0)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings2(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertError2(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

