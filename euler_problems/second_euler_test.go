package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindFibonacciSum(c *C) {
	nul_result := 0
	first_result := 1
	result := 89
	c.Assert(findFibonacciSequence(0), Equals, nul_result)
	c.Assert(findFibonacciSequence(1), Equals, first_result)
	c.Assert(findFibonacciSequence(10), Equals, result)
}

func (s *EulerSuite) Test_FindFibonacciEvenSum(c *C) {
	result := 34
	c.Assert(findFibonacciEvenNumbers(10), Equals, result)
}
