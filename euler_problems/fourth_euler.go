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
	s := ""
	revInt := []string{reverse(strconv.Itoa(n))}
	for _, x := range revInt {
		s += x
	}
	str, _ := strconv.Atoi(s)
	return str
}

func findLargestPalindrome(n int) int {
	factors := [2]int{}
	palindrome := makePalindrome(n)
	for i := 999; i > 99; i-- {
		if (palindrome/i) > 999 || i*i < palindrome {
			break
		}
		if palindrome%i == 0 {
			factors[0] = palindrome / i
			factors[1] = i
			break
		}
	}
	return palindrome
}
