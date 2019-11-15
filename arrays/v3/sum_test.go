package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("collections of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

		numFloat := []float64{1.0,2.0,3.0}
		got2 := sumFloat(numFloat)
		want2 := 6.0

		if got2 != want2{
			t.Errorf("got %v want %v given,%v",got2,want2,numFloat)
		}

	})

}
