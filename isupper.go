package piscine

func IsLower(str string) bool {
	bukva := []rune(str)
	counter := 0
	for _, letter := range bukva {
		if check(letter) {
			counter++
		}
	}
	if counter == check_len(str) {
		return true
	}
	return false
}

func check(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func check_len(str string) int {
	var count int
	strAsByte := []rune(str)
	for index := range strAsByte {
		count = index + 1
	}
	return count
}
