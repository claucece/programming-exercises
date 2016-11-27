package euler

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type EulerSuite struct{}

var _ = Suite(&EulerSuite{})

func (s *EulerSuite) TestDivisorsByBruteForce(c *C) {
	result := 233168
	c.Assert(bruteSumMultiples(1000), Equals, result)
}

func (s *EulerSuite) TestDivisorsByProgression(c *C) {
	result := 233168
	c.Assert(sumMultiples(), Equals, result)
}
