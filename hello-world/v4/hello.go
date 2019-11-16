package main

import "fmt"

const englishHelloPrefix = "Hello, "
const chineseHelloPrefix = "你好, "
// Hello returns a personalised greeting
func Hello(name string) string {
	return englishHelloPrefix + name
}

func Nihao(name string) string{
	return chineseHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
	fmt.Println(Nihao("GO"))
}
