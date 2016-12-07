package math_problems

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MathSuite struct{}

var _ = Suite(&MathSuite{})

func (s *MathSuite) Test_SwitchToArabicNumber(c *C) {
	oneCharResult := 1
	multipleCharResult := 6
	c.Assert(switchToArabic("I"), Equals, oneCharResult)
	c.Assert(switchToArabic("VI"), Equals, multipleCharResult)
}
