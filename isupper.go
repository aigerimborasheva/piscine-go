package piscine

func IsUpper(str string) bool {
	bukva := []rune(str)
	counter := 0
	for _, letter := range bukva {
		if checkUpp(letter) {
			counter++
		}
	}
	if counter == checkUpp_len(str) {
		return true
	}
	return false
}

func checkUpp(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func checkUpp_len(str string) int {
	var count int
	strAsByte := []rune(str)
	for index := range strAsByte {
		count = index + 1
	}
	return count
}
