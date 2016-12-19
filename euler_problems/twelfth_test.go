package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindNumberOfDivisors(c *C) {
	rslt := 2

	c.Assert(findDivisors(5), Equals, rslt)
}

func (s *EulerSuite) Test_FindValueOfTriangleOver500Divisor(c *C) {
	rslt := 76576500

	c.Assert(findTriangleWithNDivisors(500), Equals, rslt)
}
