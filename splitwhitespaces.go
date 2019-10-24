package piscine

func SplitWhiteSpaces(str string) []string {
	runes := []rune(str)
	count := 0

	for _, sign := range runes {
		if sign == 32 || sign == 10 || sign == 9 {
			count++
		}

	}

	k := 0
	size := count + 3
	answer := make([]string, size)

	stringe := ""
	for _, sign := range runes {
		if sign == 32 || sign == 10 || sign == 9 {
			answer[k] = stringe
			k++
			stringe = ""
		} else {
			stringe = stringe + string(sign)
		}
	}

	return answer
}
