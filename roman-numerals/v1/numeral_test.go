package v1

import "testing"

func TestRomanNumerals(t *testing.T) {
	got := ConvertToRoman(1)
	want := "I"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func ConvertToRoman(arabic int) string {
	if arabic == 1{
		return "I"
	}
	return "I"
}
