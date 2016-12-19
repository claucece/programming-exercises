package euler

import (
	"math"
)

// XXX: improve this with prime factorization and coprimes
func findDivisors(n int) int {
	nd := 0
	sqrt := math.Sqrt(float64(n))

	for i := 1; i <= int(sqrt); i++ {
		if n%i == 0 {
			nd += 2
		}
	}
	if int(sqrt)*int(sqrt) == n {
		nd--
	}
	return nd
}

func findTriangleWithNDivisors(dvr int) int {
	var valOfT int
	for i := 1; findDivisors(valOfT) <= dvr; i++ {
		valOfT += i
	}

	return valOfT
}
