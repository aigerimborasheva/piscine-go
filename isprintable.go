package piscine

func IsPrintable(str string) bool {
	bukva := []rune(str)
	for _, letter := range bukva {
		if checkPrintable(letter) {
			return false
		}
	}
	return true
}

func checkPrintable(r rune) bool {
	if r >= 0 && r <= 31 {
		return true
	}
	return false
}
