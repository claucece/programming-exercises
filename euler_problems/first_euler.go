package euler

// common brute force one
func bruteSumMultiples(n int) int {
	var result int
	for i := 1; i < n; i++ {
		if (i%3) == 0 || (i%5) == 0 {
			result += i
		}
	}
	return result
}

// By this: http://www.wikihow.com/Sum-the-Integers-from-1-to-N
// and https://en.wikipedia.org/wiki/Arithmetic_progression
func divisibleBy(dvsr, lmt int) int {
	return dvsr * (lmt / dvsr) * ((lmt / dvsr) + 1) / 2
}

func sumMultiples() int {
	return divisibleBy(3, 999) + divisibleBy(5, 999) - divisibleBy(15, 999)
}
