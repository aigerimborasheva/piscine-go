package main
import "github.com/01-edu/z01"
import "os"

func main(){
	arguments:= os.Args
	str:=arguments[0]
	bukva := []rune(str)
	for _, runes:= range bukva{
		z01.PrintRune(runes)
	}
	z01.PrintRune(10)
}
