package main

import "github.com/01-edu/z01"
import "os"

func main() {
	arguments := os.Args
	strings := []string(arguments)
	reversed := ""
	size := len(arguments)

	for _, _ = range strings {
		if size != 0 {
			reversed = strings[size]
			for _, strs := range reversed {
				z01.PrintRune(strs)
			}
			z01.PrintRune(10)
			size = size - 1
		}
	}
}
