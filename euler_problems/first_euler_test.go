package euler

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type EulerSuite struct{}

var _ = Suite(&EulerSuite{})

func (s *EulerSuite) Test_SumMultiplesByBruteForce(c *C) {
	rslt := 233168
	c.Assert(sumMultiplesByBruteForce(1000), Equals, rslt)
	c.Assert(sumMultiples(1000), Equals, rslt)
}

func (s *EulerSuite) Test_SumMultiplesByProgression(c *C) {
	rslt := 233168
	c.Assert(sumMultiples(1000), Equals, rslt)
}
