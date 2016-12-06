package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindSumOfPrimes(c *C) {
	expected_result := 17

	c.Assert(findSumOfPrimes(10), Equals, expected_result)
}
