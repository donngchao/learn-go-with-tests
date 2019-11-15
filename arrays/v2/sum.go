package main

// Sum calculates the total from an array of numbers
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func Sum2(numbersTest [6]int) int{
	sum2 := 0
	for _,number := range numbersTest{
		sum2 += number
	}
	return sum2
}
