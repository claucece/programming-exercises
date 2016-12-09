package euler

func findFibonacciSequence(n int) int {
	result := 0
	for result < n {
		switch n {
		case 0:
			return 0
		case 1:
			return 1
		default:
			numbers := findFibonacciSequence(n-1) + findFibonacciSequence(n-2)
			result += numbers
		}
	}
	return result
}

// Only even numbers
func findFibonacciEvenNumbers(n int) int {
	result, fib3 := 2, 2
	sum, fib6 := 0, 0

	for result < n {
		sum += result
		result = 4*fib3 + fib6
		fib6 = fib3
		fib3 = result
	}

	return result
}
