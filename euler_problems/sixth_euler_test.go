package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindDifferenceOfSquares(c *C) {
	rslt := 25164150

	c.Assert(difference(), Equals, rslt)
	c.Assert(calculateDifferenceOfSquares(), Equals, rslt)
}
