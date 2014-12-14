package prime

func Prime(num int) []int {
	if num < 2 {
		return []int{}
	}

	if isPrime(num) {
		return []int{num}
	}

	return []int{2, 2}
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
