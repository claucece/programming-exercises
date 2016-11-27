package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) TestFindPrimes(c *C) {
	expected_result := 104743

	c.Assert(findPrimes(10001), Equals, expected_result)
}
