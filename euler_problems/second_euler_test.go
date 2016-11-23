package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) TestFindFibonacciSum(c *C) {
	nul_result := 0
	first_result := 1
	result := 89
	c.Assert(fibonacci(0), Equals, nul_result)
	c.Assert(fibonacci(1), Equals, first_result)
	c.Assert(fibonacci(10), Equals, result)
}
