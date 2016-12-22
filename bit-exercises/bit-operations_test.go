package bits

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type BiteSuite struct{}

var _ = Suite(&BiteSuite{})

func (s *BiteSuite) Test_ComputeSignofInt(c *C) {

	n := 1
	m := -1

	c.Assert(computeSign(n), Equals, 0)
	c.Assert(computeSign(m), Equals, -1)
}
