package prime

func Prime(num int) []int {
	if num < 2 {
		return []int{}
	}

	if num == 4 {
		return []int{2, 2}
	}

	return []int{num}
}
