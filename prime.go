package prime

func Prime(number int) []int {
	if number < 2 {
		return []int{}
	}

	list := []int{}
	divisor := 2

	for number >= divisor {
		if number == 2 {
			return append(list, 2)
		}

		mod := number % divisor
		div := number / divisor
		if mod == 0 {
			list = append(list, divisor)
			number = div
		} else {
			if isPrime(number) {
				list = append(list, number)
				number = divisor
			}
			divisor++
		}

	}
	return list
}

func isPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
