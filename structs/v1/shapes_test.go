package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

	gotArea := CalculateArea(10,10)
	wantArea := 100.0
	if gotArea != wantArea {
		t.Errorf("got %.2f want %.2f",gotArea,wantArea)
	}
}
