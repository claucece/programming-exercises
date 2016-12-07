package math_problems

import (
	"strconv"
)

// TODO: try thrid case with "IV" and so on
func switchToArabic(s string) int {
	numbers := [7][7]string{{"I", "1"}, {"V", "5"}, {"X", "10"},
		{"L", "50"}, {"C", "100"}, {"D", "500"}, {"M", "1000"}}
	var sum int

	for x, _ := range numbers {
		for _, i := range s {
			if numbers[x][0] == string(i) {
				n, _ := strconv.Atoi(numbers[x][1])
				sum += n
			}
		}
	}
	return sum
}
