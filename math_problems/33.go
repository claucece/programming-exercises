package math_problems

import (
	"fmt"
	_ "strconv"
)

// TODO: try thrid case with "IV" and so on
func switchToArabic(s string) int {
	numbers := [7][7]string{{"I", "1"}, {"V", "5"}, {"X", "10"},
		{"L", "50"}, {"C", "100"}, {"D", "500"}, {"M", "1000"}}
	intNum, prev, sum := 0

	for i := len(s) - 1; i >= 0; i-- {
		for x, _ := range numbers {
			if numbers[x][0] == string(s[i]) {
				fmt.Println("bla")
			}
		}

		// for _, i := range s {
		//	if numbers[x][0] == string(i) {
		//		n, _ := strconv.Atoi(numbers[x][1])
		//		sum += n
		//	}
		//}
	}
	return sum
}
