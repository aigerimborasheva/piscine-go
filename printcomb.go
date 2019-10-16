package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	i := '0'
	k := '1'
	j := '2'

	for i <= '9' {
		for k <= '9' {
			for j <= '9' {
				if i < k && k < j {
					z01.PrintRune(i)
					z01.PrintRune(k)
					z01.PrintRune(j)
					if i != '7' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				i++
			}
			j = '0'
			k++
		}
		k = '0'
		i++
	}
	z01.PrintRune('\n')
}
