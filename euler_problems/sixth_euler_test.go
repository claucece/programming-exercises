package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindDifferenceOfSquares(c *C) {
	expected_result := 25164150

	c.Assert(difference(), Equals, expected_result)
	c.Assert(calculateDifferenceOfSquares(), Equals, expected_result)
}
