package piscine

func AppendRange(min, max int) []int {
	if max > min {
		size := max - min
		answer := make([]int, size)
		for i := 0; i <= size-1; i++ {
			if min < max {
				answer[i] = min
				min++
			}
		}
		return answer
	} else {
		return nil
	}
}
