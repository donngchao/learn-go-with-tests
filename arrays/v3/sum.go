package main

// Sum calculates the total from a slice of numbers
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// sumFloat calculates the total from a slice of float numbers
func sumFloat(floatNum []float64) float64{
	sum2 := 0.0
	for _,v := range floatNum{
		sum2 += v
	}
	return sum2
}
