package main

import "fmt"

const englishHelloPrefix = "Hello, "
const hiPrefix  = "Hi, "

// Hello returns a personalised greeting, defaulting to Hello, world if an empty name is passed
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}
//reverse logic
func Hi(name string) string{
	if name != ""{
		name = "World"
	}else{
		name = "GO"
	}
	return hiPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
	fmt.Println(Hi(""))//Hi, GO
	fmt.Println(Hi("there"))//Hi, World

}
