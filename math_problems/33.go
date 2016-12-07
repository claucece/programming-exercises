package math_problems

import "strconv"

// TODO: try second case with "II" and so on
func switchToArabic(s string) int {
	numbers := [7][7]string{{"I", "1"}, {"V", "5"}, {"X", "10"},
		{"L", "50"}, {"C", "100"}, {"D", "500"}, {"M", "1000"}}
	var sum int

	for x, _ := range numbers {
		if numbers[x][0] == s {
			n, _ := strconv.Atoi(numbers[x][1])
			sum += n
		}
	}
	// fmt.Println(numbers[0][0])
	return sum
}
