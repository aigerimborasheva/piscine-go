package piscine 
import "github.com/01-edu/z01"
func PrintComb(){
	i := '0'
	k := '1'
	j := '2'
	for ; i <= '7'; i++ {
		for ; k <= '8'; k++ {
			for ; j <= '9'; j++ {
				if i<k && k< j {
					z01.PrintRune(i)
                    z01.PrintRune(',')
					z01.PrintRune(' ')
			}
		}
		
	}
	z01.PrintRune('\n')
}
