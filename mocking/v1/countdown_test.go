package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestCountdownVersion2(t *testing.T) {
	buffer := &bytes.Buffer{}

	CountdownVersion2(buffer)

	got := buffer.String()
	want := "4"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
