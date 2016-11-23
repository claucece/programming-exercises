package main

import (
	"fmt"
)

func findPrimes(n int) int {
	var isPrime bool
	primes := []int{2}
	counter := 1
	for len(primes) < n {
		counter += 2
		j := 0
		isPrime = true
		for primes[j]*primes[j] <= counter {
			if counter%primes[j] == 0 {
				isPrime = false
				break
			}
			j++
		}
		if isPrime {
			primes = append(primes, counter)
		}
	}
	return primes[len(primes)-1]
}

func main() {
	fmt.Println(findPrimes(1001))
}
