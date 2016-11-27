package euler

import (
	"github.com/stretchr/testify/assert"
	. "gopkg.in/check.v1"
	"testing"
)

func returnThree(a, b, c interface{}) []interface{} {
	return []interface{}{a, b, c}
}

func (s *EulerSuite) Test_GreatestCommonDivisor(c *C) {
	expected_result := 4
	c.Assert(gcd(8, 4), Equals, expected_result)
}

func Test_FindPithagoreanTriplet(t *testing.T) {

	assert.Equal(t, returnThree(findAPithagoreanTriplet()), returnThree(3, 4, 5))
	assert.Equal(t, returnThree(findAPithagoreanTripletByCondition()), returnThree(200, 375, 425))
	assert.Equal(t, returnThree(findPithagoreanTripletByEuclid()), returnThree(375, 200, 425))
}
