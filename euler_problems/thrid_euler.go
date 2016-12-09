package euler

// by this theorem: Any integer greater than 1 is either a prime number,
// or can be written as a unique product of prime numbers (ignoring the order).
func findLargestPrime(n int) int {
	lrgp, dvsr := 0, 2
	for dvsr*dvsr <= n {
		if n%dvsr == 0 {
			n /= dvsr
			lrgp = n
		}
		dvsr++
	}
	return lrgp
}

//by brute force
func findLargestPrimeByBruteForce(n int) int {
	lrgp := 0
	for i := 2; i < n; i++ {
		if n%i == 0 {
			lrgp = i
		}
	}
	return lrgp
}
