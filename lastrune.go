package piscine


func LastRune(s string) rune {
	sentence := []rune(s)
	len := Length(s)
	return sentence[len-1]
	}
