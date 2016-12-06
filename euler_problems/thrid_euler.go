package euler

// by this theorem: Any integer greater than 1 is either a prime number,
// or can be written as a unique product of prime numbers (ignoring the order).
func findLargestFactor(n int) int {
	var lrgst_fctr int
	counter := 2
	for counter*counter <= n {
		if n%counter == 0 {
			n /= counter
			lrgst_fctr = n
		}
		counter++
	}
	return lrgst_fctr
}

//by brute force
func findLargestFactorByBruteForce(n int) int {
	var lrgs_fctr int
	for i := 2; i < n; i++ {
		if n%i == 0 {
			lrgs_fctr = i
		}
	}
	return lrgs_fctr
}
