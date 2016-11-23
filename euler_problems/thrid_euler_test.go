package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) TestFindFirstPrimes(c *C) {
	result := 9

	c.Assert(findFirstPrimes(10), Equals, result)
}
