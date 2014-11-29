package prime

func Prime(number int) []int {
	if number < 2 {
		return []int{}
	}

	list := []int{}
	divisor := 2
	if number%divisor == 0 && number > 2 {
		list = append(list, divisor)
		list = append(list, number/divisor)
	} else {
		list = append(list, number)
	}

	return list
}
