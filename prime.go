package prime

func Prime(num int) []int {
	if num < 2 {
		return []int{}
	}
	return []int{num}
}
