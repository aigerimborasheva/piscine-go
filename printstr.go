package piscine
import "github.com/01-edu/z01"

func PrintStr(str string) {
	slictt := []rune(str)
	for _, str := range slictt {
		z01.PrintRune(str)
	}

	z01.PrintRune(10)
}
