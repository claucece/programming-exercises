package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) TestFindLagestPalindrome(c *C) {
	result := 899

	c.Assert(checkEquality(998), Equals, result)
}
