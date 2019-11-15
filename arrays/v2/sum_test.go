package main

import "testing"

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}

	/**
	use func Sum2()
	*/
	numbers2 := [6]int{2,2,2,3,3,3}
	got2 := Sum2(numbers2)
	want2 := 15

	if got2 != want2 {
		t.Errorf("got %d wan %d,%v",got2,want2,numbers2)
	}

}
