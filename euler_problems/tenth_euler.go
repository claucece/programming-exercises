package euler

// still in development. Will probably be solvedby Sieve. Just trying to look
// for a way to brute force it with small values.
func findSumOfPrimes(n int) int {
	var sum int
	for i := 2; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
