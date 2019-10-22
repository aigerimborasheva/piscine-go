package piscine

func ToUpper(s string) string {
	Bukva := []rune(s)
	for i := range Bukva {
		if Bukva[i] >= 'a' && Bukva[i] <= 'z' {
			Bukva[i] = Bukva[i] - 32
		}
	}
	return string(Bukva)
}
