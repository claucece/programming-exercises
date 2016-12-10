package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindiProductOfFourAdjecentNumbersInGrid(c *C) {
	expctRst := 70600674

	c.Assert(findProductOfAdjecentNumbersByGrid(4), Equals, expctRst)
	// c.Assert(findProductOfAdjecentNumbersByFile(4), Equals, expctRst)
}
