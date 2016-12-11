package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindFibonacciSum(c *C) {
	nulRslt := 0
	frstRslt := 1
	rslt := 89
	c.Assert(findFibonacciSequence(0), Equals, nulRslt)
	c.Assert(findFibonacciSequence(1), Equals, frstRslt)
	c.Assert(findFibonacciSequence(10), Equals, rslt)
}

func (s *EulerSuite) Test_FindFibonacciEvenSum(c *C) {
	rslt := 34
	c.Assert(findFibonacciEvenNumbers(10), Equals, rslt)
}

func (s *EulerSuite) Test_FindLargestFibonacci(c *C) {
	rslt := "55"
	c.Assert(findLargestFibonacciInBigInt(10), Equals, rslt)
}
