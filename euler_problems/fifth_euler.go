package euler

import "math"

func generatePrimes(lmt int) []int {
	primes := []int{2}
	for i := 3; i <= lmt; i += 2 {
		isPrime := true
		for j := 0; primes[j]*primes[j] <= i; j++ {
			if i%primes[j] == 0 {
				isPrime = false
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

// hate this casting
func findEvenlyDivisible(lmt int) int {
	primes := generatePrimes(lmt)
	posN := 1.0

	for i := 0; i < len(primes); i++ {
		frstPrimes := float64(primes[i])
		exp := math.Floor(math.Log(float64(lmt)) / math.Log(frstPrimes))
		posN = posN * (math.Pow(frstPrimes, exp))
	}
	return int(posN)
}
