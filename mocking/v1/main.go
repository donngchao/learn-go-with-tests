package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown prints a countdown from 3 to out
func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}

func CountdownVersion2(out io.Writer) {
	fmt.Fprint(out, "4")
}

func main() {
	Countdown(os.Stdout)
	CountdownVersion2(os.Stdout)
}
