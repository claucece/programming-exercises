package main

import (
	"fmt"
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
	return strconv.Atoi(rev_int)
}

func checkEquality(n int) int {
	factors := []int{}
	first_half := makePalindrome(n)
	for i := 99; i > 9; i-- {
		if (n/i) > 99 || i*i < first_half {
			break
		}
		if first_half%i == 0 {
			factors[0] = first_half / i
			factors[1] = first_half
		}
	}
}

func main() {
	fmt.Println(makePalindrome(10))
}
