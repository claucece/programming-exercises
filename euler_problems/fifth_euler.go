package euler

import (
	"strconv"
)

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func makePalindrome(n int) int {
	rev_int := []string{reverse(strconv.Itoa(n))}
	var s string
	for _, x := range rev_int {
		s += x
	}
	str, _ := strconv.Atoi(s)
	return str
}

func checkEquality(n int) int {
	n--
	factors := [2]int{}
	first_half := makePalindrome(n)
	for i := 999; i > 99; i-- {
		if (first_half/i) > 999 || i*i < first_half {
			break
		}
		if first_half%i == 0 {
			factors[0] = first_half / i
			factors[1] = first_half
			break
		}
	}
	return first_half
}
