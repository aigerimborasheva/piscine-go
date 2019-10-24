package piscine

func ConcatParams(args []string) string {
	slova := []string(args)
	answers := ""
	for index, slov := range slova {
		if index > 0 {
			answers = answers + "\n" + slov
		} else {
			answers = answers + slov
		}
	}
	return answers
}
