package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindLagestPalindrome(c *C) {
	rslt := 899

	c.Assert(findLargestPalindrome(998), Equals, rslt)
}
