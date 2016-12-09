package euler

import (
	"github.com/bradfitz/iter"
)

// common brute force one
func sumMultiplesByBruteForce(n int) int {
	result := 0
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}
	return result
}

// By this: http://www.wikihow.com/Sum-the-Integers-from-1-to-N
// and https://en.wikipedia.org/wiki/Arithmetic_progression
func divisibleBy(dvsr, lmt int) int {
	return dvsr * (lmt / dvsr) * (lmt/dvsr + 1) / 2
}

func sumMultiplesByProgression() int {
	return divisibleBy(3, 999) + divisibleBy(5, 999) - divisibleBy(15, 999)
}

// interesting but pretty useless
func sumMultiples(n int) int {
	result := 0
	for x := range iter.N(n) {
		if x%3 == 0 || x%5 == 0 {
			result += x
		}
	}
	return result
}
