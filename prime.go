package prime

func Prime(num int) []int {
	if num < 2 {
		return []int{}
	}

	if isPrime(num) {
		return []int{num}
	}

	result := []int{}
	div := 2
	mod := num % div
	if mod == 0 {
		result = append(result, div)
		result = append(result, num/div)
	}

	return result
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
