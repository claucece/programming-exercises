package euler

import "math"

func generatePrimes(lmt int) []int {
	primes := []int{2}
	var isPrime bool

	for i := 3; i <= lmt; i += 2 {
		j := 0
		isPrime = true
		for primes[j]*primes[j] <= i {
			if i%primes[j] == 0 {
				isPrime = false
				break
			}
			j++
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

func findSmallerDivisor(lmt int) int {
	primes := generatePrimes(lmt)
	result := 1.0

	for i := 0; i < len(primes); i++ {
		a := math.Floor(math.Log(float64(lmt)) / math.Log(float64(primes[i])))
		result = result * (math.Pow(float64(primes[i]), a))
	}
	return int(result)
}
