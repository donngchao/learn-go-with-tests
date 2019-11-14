package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got := dictionary.Search("test")
	want := "this is just a test"

	assertStrings(t, got, want)

	telephoneBook := Telephone{"Jack":13511112222}

	phoneNum := telephoneBook.Search("Jack")
	wantNum := 13511112222
	assertInt(t,phoneNum,wantNum)

	priceList := Price{1:1,2:2}

	priceNow := priceList.GetPrice(2)
	wantPrice := 2
	assertInt(t,priceNow,wantPrice)

}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertInt(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
