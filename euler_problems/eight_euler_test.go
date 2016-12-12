package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindGreastestProduct(c *C) {
	rslt := 23514624000

	c.Assert(findLargestProductOf(13), Equals, rslt)
}
