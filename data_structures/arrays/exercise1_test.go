package main

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type ArraySuite struct{}

var _ = Suite(&ArraySuite{})

func (s *ArraySuite) Test_Average(c *C) {
	result := 150
	total := []int{100, 200}
	c.Assert(findAverage(total), Equals, result)
}

func (s *ArraySuite) Test_FindString(c *C) {
	result := 3
	str := "lololo"
	c.Assert(findString(str, "l"), Equals, result)
}
