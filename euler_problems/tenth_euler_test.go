package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindSumOfPrimes(c *C) {
	expected_result := 17

	c.Assert(findSumOfPrimesByBruteForce(10), Equals, expected_result)
	c.Assert(findPrimesBySieve(10), Equals, expected_result)
}
