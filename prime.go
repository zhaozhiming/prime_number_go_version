package prime

func Prime(num int) []int {
	if num < 2 {
		return []int{}
	}

	result := []int{}
	for div := 2; div <= num; {
		if num%div == 0 {
			result = append(result, div)
			num /= div
		} else {
			div++
		}
	}
	return result
}
