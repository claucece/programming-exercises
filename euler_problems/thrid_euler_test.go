package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindLargestPrime(c *C) {
	rslt := 5
	lgRslt := 6857

	c.Assert(findLargestPrimeByBruteForce(10), Equals, rslt)
	c.Assert(findLargestPrime(600851475143), Equals, lgRslt)
}
