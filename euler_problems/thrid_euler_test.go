package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) TestFindFirstPrimes(c *C) {
	result := 5
	secondResult := 6857

	c.Assert(findLargestFactorByBruteForce(10), Equals, result)
	c.Assert(findLargestFactor(600851475143), Equals, secondResult)
}
