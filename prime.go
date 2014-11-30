package prime

func Prime(number int) []int {
	if number < 2 {
		return []int{}
	}

	list := []int{}

	for divisor := 2; number >= divisor; {
		if isPrime(number) {
			return append(list, number)
		}

		mod := number % divisor
		div := number / divisor
		if mod == 0 {
			list = append(list, divisor)
			number = div
		} else {
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
