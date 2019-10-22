package piscine

func ToLower(s string) string {
	bukvalow := []byte(s)
	for index, letter := range bukvalow {
		if letter <= 90 && letter >= 65 {
			bukvalow[index] = letter + 32
		}
	}
	return string(bukvalow)
}
