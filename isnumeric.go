package piscine

func IsNumeric(str string) bool {
	bucva := []rune(str)
	counter := 0
	for _, letter := range bucva {
		if checkNumber(letter) {
			counter++
		}
	}
	if counter == StrLenNum(str) {
		return true
	}
	return false
}

func checkNumber(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func StrLenNum(str string) int {
	var count int
	strAsByte := []rune(str)
	for index := range strAsByte {
		count = index + 1
	}
	return count
}
