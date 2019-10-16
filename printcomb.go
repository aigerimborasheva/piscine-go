package piscine 
import "github.com/01-edu/z01"
func PrintComb(){
	var i rune = '0'
	var k rune = '1'
	var j rune = '2'
	for i <= '7'; i++ {
		for k <= '8'; k++ {
			for j <= '9'; j++ {
				if i<k && k< j {
					z01.PrintRune(i)
                    z01.PrintRune(',')
					z01.PrintRune(' ')
			}
		}
		
	}
	z01.PrintRune('\n')
}
