package euler

// TODO: altough sieve is fine, there is also Atkin.

// a weird brute force one. It is only working until 25, of course.
func findSumOfPrimesByBruteForce(n int) int {
	sum := 0
	for i := 2; i < n; i++ {
		if i == 2 || i%2 != 0 {
			if i == 3 || i%3 != 0 {
				sum += i
			}
		}
	}
	return sum
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func sumPrimes(primes []int) int {
	sum := 0
	for _, i := range primes {
		sum += i
	}
	return sum
}

// TODO: check the math of this
func findPrimesBySieve(upperN int) int {
	primes := []int{2}
	notPrimes := []int{1}

	for i := 3; i <= upperN; i += 2 {
		for j := i * i; j <= upperN; j += 2 * i {
			notPrimes = append(notPrimes, j)
		}
		if !(contains(notPrimes, i)) {
			primes = append(primes, i)
		}
	}
	return sumPrimes(primes)
}
