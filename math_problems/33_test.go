package math_problems

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MathSuite struct{}

var _ = Suite(&MathSuite{})

func (s *MathSuite) Test_SwitchToArabicNumber(c *C) {
	result := 1
	c.Assert(switchToArabic("I"), Equals, result)
}
