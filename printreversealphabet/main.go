package main

import "github.com/01-edu/z01"

func main() {
	var aVarRune rune = 'z'
	for aVarRune >= 'a' {
		z01.PrintRune(aVarRune)
		aVarRune--
	}
	z01.PrintRune(10)
}
