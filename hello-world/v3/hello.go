package main

import "fmt"

// Hello returns a personalised greeting
func Hello(name string) string {
	return "Hello, " + name
}

func Hi(name string) string{
	return "Hi, " + name
}

func main() {
	fmt.Println(Hello("world"))
	fmt.Println(Hi("Mike"))
}
