package piscine

func AppendRange(min, max int) []int {
	if max > min {
		size := max - min
		var answer []int
		for i := 0; i <= size-1; i++ {
			if min < max {
				answer = append(answer, min)
				min++
			}
		}
		return answer
	} else {
		return nil
	}
}
