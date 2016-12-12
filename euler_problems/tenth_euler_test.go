package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindSumOfPrimes(c *C) {
	rslt := 17

	c.Assert(findSumOfPrimesByBruteForce(10), Equals, rslt)
	c.Assert(findPrimesBySieve(10), Equals, rslt)
}
