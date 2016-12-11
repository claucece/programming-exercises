package euler

import (
	"math/big"
)

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

// with big numbers
func findLargestFibonacciInBigInt(n int) string {
	last := big.NewInt(1)
	current := big.NewInt(1)

	for i := 2; i < n; i++ {
		last.Add(last, current)
		tmp := last
		last = current
		current = tmp
	}
	return current.String()
}
