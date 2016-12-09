package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindLargestPrime(c *C) {
	result := 5
	longResult := 6857

	c.Assert(findLargestPrimeByBruteForce(10), Equals, result)
	c.Assert(findLargestPrime(600851475143), Equals, longResult)
}
