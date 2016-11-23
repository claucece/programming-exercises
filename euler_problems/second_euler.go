package euler

func fibonacci(n int) int {
	var result int
	for result < n {
		switch n {
		case 0:
			return 0
		case 1:
			return 1
		default:
			numbers := fibonacci(n-1) + fibonacci(n-2)
			result += numbers
		}
	}
	return result
}
