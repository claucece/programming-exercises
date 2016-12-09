package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindSmallestEvenlyDivisibleNumber(c *C) {
	result := 232792560

	c.Assert(findEvenlyDivisible(20), Equals, result)
}
