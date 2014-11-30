package prime

func Prime(number int) []int {
	if number < 2 {
		return []int{}
	}

	list := []int{}
	divisor := 2

	for number >= divisor {
		mod := number % divisor
		div := number / divisor
		if mod == 0 && number > 2 {
			list = append(list, divisor)
		} else {
			list = append(list, number)
			divisor++
		}
		number = div
	}

	return list
}
