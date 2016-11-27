package euler

import (
	"math"
)

// by brute force. Traditional Pythagorean triplet
func findAPithagoreanTriplet() (int, int, int) {
	var a, b, c int
	for c = 0; c <= 1000; c++ {
		for b = 0; b < c; b++ {
			for a = 0; a < b; a++ {
				if a*a+b*b == c*c {
					return a, b, c
				}

			}
		}
	}
	return 0, 0, 0
}

// by brute force but also applying this condition: a < 1000/3,  and a < b < 1000/2
func findAPithagoreanTripletByCondition() (int, int, int) {
	var a, b, c int
	s := 1000
	for a = 1; a <= s/3; a++ {
		for b = a; b < s/2; b++ {
			c = s - a - b
			if (a*a)+(b*b) == (c * c) {
				return a, b, c
			}

		}
	}
	return 0, 0, 0
}

// number theoretical approach: euclides formula and primitive(a A Pythagorean Triplet is
// called primitive if a, b and c are coprime. That means that the greatest common divisor (gcd) of
// the set {a,b,c} is one.

// find the gcd
func gcd(a, b int) int {
	var x, y int

	if a > b {
		x = a
		y = b
	} else {
		x = b
		y = a
	}

	for x%y != 0 {
		var temp int
		temp = x
		x = y
		y = temp % x
	}
	return y
}

func findPithagoreanTripletByEuclid() (int, int, int) {
	var a, b, c int
	s := 1000
	var m, k, n, d int
	found := false

	limit := int(math.Sqrt(float64(s) / 2.0))

	for m = 2; m <= limit; m++ {
		if (s/2)%m == 0 {
			if m%2 == 0 {
				k = m + 1
			} else {
				k = m + 2
			}
			for k < 2*m && k <= s/(2*m) {
				if s/(2*m)%k == 0 && gcd(k, m) == 1 {
					d = s / 2 / (k * m)
					n = k - m
					a = d * (m*m - n*n)
					b = 2 * d * n * m
					c = d * (m*m + n*n)
					found = true
					break
				}
				k += 2
			}
		}
		if found {
			break
		}
	}
	return a, b, c
}
