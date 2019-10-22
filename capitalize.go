package piscine

func Capitalize(s string) string {
	bukva:= []rune(s)
	for index, letter := range bukva{
		if checkAlphNum(letter) {
			if index == 0 || checkAlphNum(stringAsRune[index-1]) == false {
				if letter >= 'a' && letter <= 'z' {
					tring[index] = letter - 32
				}
			} else {
				if letter >= 'A' && letter <= 'Z' {
					tring[index] = letter + 32
				}
			}
		}
	}
	return string(tring)
}

func checkAlphNum(r rune) bool {
	if r >= 'A' && r <= 'Z' ||
		r >= 'a' && r <= 'z' ||
		r >= '0' && r <= '9' {
		return true
	}
	return false
}