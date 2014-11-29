package prime

func Prime(number int) []int {
	if number < 2 {
		return []int{}
	}

	return []int{number}
}
