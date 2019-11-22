package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collections of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllAgain(t *testing.T) {

	got := SumAll([]int{17, 8}, []int{1, 3,7,1,0})
	want := []int{25, 12}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
