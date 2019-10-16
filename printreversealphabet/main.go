package main

import "github.com/01-edu/z01"

func main() {
	var l rune
	l := 122
	for l > 96 {
		z01.PrintRune(l)
		l--
	}
	z01.PrintRune(10)
}
