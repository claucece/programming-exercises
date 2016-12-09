package euler

func findPrime(n int) int {
	primes := []int{2}
	smallp := 1
	for len(primes) < n {
		smallp += 2
		isPrime := true
		for j := 0; primes[j]*primes[j] <= smallp; j++ {
			if smallp%primes[j] == 0 {
				isPrime = false
			}
		}
		if isPrime {
			primes = append(primes, smallp)
		}
	}
	return primes[len(primes)-1]
}
