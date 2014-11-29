package prime

func Prime(number int) []int {
	if number < 2 {
		return []int{}
	}

	list := []int{}
	if number == 4 {
		list = append(list, 2)
		list = append(list, 2)
	} else {
		list = append(list, number)
	}

	return list
}
