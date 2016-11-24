package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) TestSmallerDivisor(c *C) {
	expected_result := 232792560

	c.Assert(findSmallerDivisor(20), Equals, expected_result)
}
