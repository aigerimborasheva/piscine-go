package piscine

func Capitalize(s string) string {
	bukvy := []rune(s)
	for index, letter := range bukvy {
		if checkAlphNum(letter) {
			if index == 0 || checkAlphNum(bukvy[index-1]) == false {
				if letter >= 'a' && letter <= 'z' {
					bukvy[index] = letter - 32
				}
			} else {
				if letter >= 'A' && letter <= 'Z' {
					bukvy[index] = letter + 32
				}
			}
		}
	}
	return string(bukvy)
}

func checkAlphNum(r rune) bool {
	if r >= 'A' && r <= 'Z' ||
		r >= 'a' && r <= 'z' ||
		r >= '0' && r <= '9' {
		return true
	}
	return false
}
