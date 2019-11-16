package main

import "fmt"

// Hello returns a greeting
func Hello() string {
	return "Hello, world"
}

func Hi() string{
	return "Hi,there! I am using golang"
}

func main() {
	fmt.Println(Hello())
	fmt.Println(Hi())
}
