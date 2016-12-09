package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindNPrime(c *C) {
	expected_result := 104743

	c.Assert(findPrime(10001), Equals, expected_result)
}
