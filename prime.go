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

		if number%divisor != 0 {
			divisor++
			continue
		}

		list = append(list, divisor)
		number = number / divisor
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
