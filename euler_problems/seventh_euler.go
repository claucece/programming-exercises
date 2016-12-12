package euler

func findPrime(n int) int {
	primes := []int{2}
	smallPrm := 1
	for len(primes) < n {
		smallPrm += 2
		isPrime := true
		for j := 0; primes[j]*primes[j] <= smallPrm; j++ {
			if smallPrm%primes[j] == 0 {
				isPrime = false
			}
		}
		if isPrime {
			primes = append(primes, smallPrm)
		}
	}
	return primes[len(primes)-1]
}
