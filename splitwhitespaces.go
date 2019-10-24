package piscine

func SplitWhiteSpaces(str string) []string {
	runes := []rune(str)
	count := 0
	mor := 0
	for _, sign := range runes {
		if sign == 32 || sign == 10 || sign == 9 {
			count++
		}
		mor++
	}

	k := 0
	size := count + 3
	answer := make([]string, size)

	stringe := ""
	for index, sign := range runes {
		if sign == 32 || sign == 10 || sign == 9 || index == mor {
			answer[k] = stringe
			k++
			stringe = ""
		} else {
			stringe = stringe + string(sign)
		}
	}

	return answer
}
