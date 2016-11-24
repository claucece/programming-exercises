package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) TestFindDifference(c *C) {
	expected_result := 25164150

	c.Assert(difference(100, 100), Equals, expected_result)
	c.Assert(calculateDifference(), Equals, expected_result)
}
