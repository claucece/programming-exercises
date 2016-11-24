package euler

// by this theorem: Any integer greater than 1 is either a prime number,
// or can be written as a unique product of prime numbers (ignoring the order).
func findLargestFactor(num int) int {
	counter := 2
	var largest_fact int
	for counter*counter <= num {
		if num%counter == 0 {
			num /= counter
			largest_fact = counter
		} else {
			counter++
		}
	}
	if num > 0 {
		largest_fact = num
	}
	return largest_fact
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
