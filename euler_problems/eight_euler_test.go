package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindGreastestProduct(c *C) {
	expected_result := 23514624000

	c.Assert(findLargestProductOfThridteen(), Equals, expected_result)
}
