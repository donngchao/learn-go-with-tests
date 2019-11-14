package main

import "testing"

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}

	numbers1 := [6]int{1,2,3,4,5,6}
	got1 := Sum1(numbers1)
	want1 := 21
	if want != got{
		t.Errorf("got %d want %d,%v",got1,want1,numbers1)
	}
}
