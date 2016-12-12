package euler

import (
	. "gopkg.in/check.v1"
)

func (s *EulerSuite) Test_FindProductOfFourAdjecentNumbersInGrid(c *C) {
	rslt := 70600674

	c.Assert(findProductOfAdjecentNumbersByGrid(4), Equals, rslt)
	// c.Assert(findProductOfAdjecentNumbersByFile(4), Equals, rslt)
}
