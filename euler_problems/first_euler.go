package euler

// common brute force one
func bruteSumMultiples(n int) int {
	result := 0
	for i := 1; i < n; i++ {
		if ((i % 3) == 0) || ((i % 5) == 0) {
			result += i
		}
	}
	return result
}

// By this: http://www.wikihow.com/Sum-the-Integers-from-1-to-N
// and https://en.wikipedia.org/wiki/Arithmetic_progression
func sumDivisibleBy(dvsr, lmt int) int {
	return dvsr * (lmt / dvsr) * ((lmt / dvsr) + 1) / 2
}

func sumMultiples() int {
	return sumDivisibleBy(3, 999) + sumDivisibleBy(5, 999) - sumDivisibleBy(15, 999)
}
